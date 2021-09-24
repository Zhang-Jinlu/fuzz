package fuzz

import "testing"

func TestFuzz(t *testing.T) {
	f := New()
	f.GET("/", func(c *Context) {

	})
	rg := f.Group("v1")
	{
		rg.GET("/hello/:name", func(c *Context) {
			c.String(200, "Hello, %s\n", c.Param("name"))
		})
	}
	f.Run(":18080")
}