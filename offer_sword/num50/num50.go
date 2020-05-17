package main

import "fmt"

// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
func firstUniqChar(s string) byte {
	mark := make(map[byte]int, len(s))
	for _, ch := range s {
		mark[byte(ch)]++
	}
	for _, ch := range s {
		if mark[byte(ch)] == 1 {
			return byte(ch)
		}
	}
	return ' '
}

func main() {
	fmt.Printf("%s\n", string(firstUniqChar("abaccdeff")))
	fmt.Printf("%s\n", string(firstUniqChar("")))
}