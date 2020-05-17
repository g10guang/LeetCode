package main

import "fmt"

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	start, end := 0, 0
	mark := make(map[byte]bool)
	for ;end < len(s); end++ {
		char := s[end]
		if mark[char] {
			// find duplicate char
			for ;s[start] != char; start++ {
				delete(mark, s[start])
			}
			delete(mark, s[start])
			start++
		}
		mark[s[end]] = true
		if end - start + 1 > maxLen {
			maxLen = end - start + 1
		}
	}
	return maxLen
}

func main() {
	fmt.Printf("%v\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("pwwkew"))
}