package fuzz

import "strings"

func parsePattern(pattern string) []string {
	if pattern == "/" {
		return []string{"/"}
	}
	split := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, s := range split {
		if s != "" {
			parts = append(parts, s)
			if s[0] == '*' {
				break
			}
		}
	}
	return parts
}
