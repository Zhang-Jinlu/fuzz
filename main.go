package main

import (
	"fmt"
	"net/http"
)

func main() {
	fuzz := New()
	fuzz.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "query url: %s\n", r.URL)
	})
	fuzz.Run(":8888")
}