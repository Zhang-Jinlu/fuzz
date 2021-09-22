package fuzz

import (
	"errors"
)

type node struct {
	part     string           // 路径中的一部分，例如 :lang; :xxx 匹配一个参数，*xxx 匹配多个参数，*只能在最后一部分
	pattern  string           // 完整路径，例如 /p/:lang 只存在于结尾的节点
	children map[string]*node // 子节点，例如 [doc, tutorial, intro]
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{
		root: &node{
			children: make(map[string]*node),
		},
	}
}

func (t *trie) insert(path string) {
	if path == "" {
		panic(errors.New("can not add empty path"))
	}
	parts := parsePattern(path)
	cur := t.root
	for _, part := range parts {
		if _, ok := cur.children[part]; !ok {
			cur.children[part] = &node{part: part, children: make(map[string]*node)}
		}
		cur = cur.children[part]
	}
	cur.pattern = path
}

func (t *trie) search(path string) *node {
	if path == "" {
		panic(errors.New("no empty path"))
	}
	parts := parsePattern(path)
	cur := t.root
	for _, part := range parts {
		nextPart := ""
		if _, ok := cur.children[part]; ok {
			nextPart = part
		} else{
			for k, _ := range cur.children {
				if k[0] == '*' || k[0] == ':' {
					nextPart = k
					break
				}
			}
		}
		if nextPart == "" {
			if cur.part[0] == '*' {
				break
			}
			return nil
		}
		cur = cur.children[nextPart]
	}
	if cur.pattern == "" {
		return nil
	}
	return cur
}