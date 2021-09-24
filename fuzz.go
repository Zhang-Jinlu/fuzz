package fuzz

import (
	"net/http"
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

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
