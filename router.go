package fuzz

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type router struct {
	roots    map[string]*trie
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*trie),
		handlers: make(map[string]HandlerFunc),
	}
}

// 添加路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = newTrie()
	}
	r.roots[method].insert(pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 获取路由 与 参数
func (r *router) getRoute(method string, pattern string) (*node, map[string]string) {
	searchParts := parsePattern(pattern)
	params := make(map[string]string)
	t, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := t.search(searchParts)
	if n != nil {
		parts := parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			} else if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
		return
	}
	c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
}
