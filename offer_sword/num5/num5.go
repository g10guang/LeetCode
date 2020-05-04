package main
import (
	"strings"
)

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	builder := strings.Builder{}
	length := 0
	for _, ch := range s {
		if ch == ' ' {
			length+=3
		} else {
			length++
		}
	}
	builder.Grow(length)
	for _, ch := range s {
		if ch == ' ' {
			builder.WriteString("%20")
		} else {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}