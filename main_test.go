package fuzz

import (
	"log"
	"testing"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func TestFuzz(t *testing.T) {
	f := New()
	f.Use(Logger())
	f.GET("/hello/:name", func(c *Context) {
		c.String(200, "Hello, %s!\n", c.Param("name"))
	})
	f.Run(":18080")
}
