package fuzz

import (
	"net/http"
	"strings"
)

// Engine 实现了Handler接口
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New fuzz.Engine的构造方法
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Run 启动Http服务的方法
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// 实现http.Handler接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, rg := range e.groups {
		if strings.HasPrefix(req.URL.Path, rg.prefix) {
			middlewares = append(middlewares, rg.middlewares...)
		}
	}
	c := newContext(w, req)
	// 将请求对应路由组中的中间件放在上下文中
	c.handlers = middlewares
	e.router.handle(c)
}
