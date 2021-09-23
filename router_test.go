package fuzz

import (
	"fmt"
	"testing"
)

func TestRouter(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/", defaultHandler)
	r.addRoute("GET", "/hello/:name", defaultHandler)
	r.addRoute("GET", "/hello/b/c", defaultHandler)
	r.addRoute("GET", "/hi/:name", defaultHandler)
	r.addRoute("GET", "/assets/*filepath", defaultHandler)

	r.handle(&Context{
		Path:   "/assets/local/redis/bin",
		Method: "GET",
	})

}

func defaultHandler(c *Context) {
	fmt.Printf("Method: %s, Path: %s, Params: %v\n", c.Method, c.Path, c.Params)
}
