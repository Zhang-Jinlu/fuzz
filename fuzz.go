package fuzz

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Engine 实现了Handler接口
type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
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
	log.Fatalf("server is listening on %s\n", addr)
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
	// 上线文中添加engine的指针
	c.engine = e
	// 将请求对应路由组中的中间件放在上下文中
	c.handlers = middlewares
	e.router.handle(c)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}
