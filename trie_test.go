package fuzz

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrie(t *testing.T) {
	tree := newTrie()
	tree.insert("/get/one/more/lesson")
	tree.insert("/add/one/more/lesson")
	tree.insert("/get/lesson/schedule")
	tree.insert("/:method/one/more/lesson")
	tree.insert("/delete/*filePath")
}

func TestOther(t *testing.T) {
	fmt.Println(strings.Split("/*", "/"))
}
