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

	fmt.Println(tree.search("/get/one/more"))
	fmt.Println(tree.search("/delete/all"))
	fmt.Println(tree.search("/update/one/more/lesson"))
	fmt.Println(tree.search("/update/more/lesson"))
}

func TestOther(t *testing.T) {
	fmt.Println(strings.Split("/*", "/"))
}
