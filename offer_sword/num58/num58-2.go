package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	builder := strings.Builder{}
	for idx := n; idx < len(s); idx++ {
		builder.WriteByte(s[idx])
	}
	for idx := 0; idx < n; idx++ {
		builder.WriteByte(s[idx])
	}
	return builder.String()
}

func main() {
	fmt.Printf("%s\n", reverseLeftWords("abcdefg", 2))
	fmt.Printf("%s\n", reverseLeftWords("lrloseumgh", 6))
}