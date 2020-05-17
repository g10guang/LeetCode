package main

import (
	"fmt"
)

var charList []byte
var result []string

// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
func permutation(s string) []string {
	result = make([]string, 0)
	charList = []byte(s)
	dfs(0)
	return result
}

func dfs(index int) {
	if index == len(charList) {
		result = append(result, string(charList))
		return
	}
	mark := make(map[byte]bool)
	for i := index; i < len(charList); i++ {
		char := charList[i]
		if mark[char] {
			continue
		}
		mark[char] = true
		swap(index, i)
		dfs(index + 1)
		swap(index, i)
	}
}

func swap(i, j int) {
	charList[i], charList[j] = charList[j], charList[i]
}

func main() {
	result := permutation("aabb")
	fmt.Printf("len=%d result=%+v\n", len(result), result)
}
