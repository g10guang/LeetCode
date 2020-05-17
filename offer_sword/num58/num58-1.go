package main

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	var tmp []string
	for _, word := range words {
		if word != "" {
			tmp = append(tmp, word)
		}
	}
	for low, high := 0, len(tmp)-1; low < high; low, high = low+1, high-1 {
		tmp[low], tmp[high] = tmp[high], tmp[low]
	}
	return strings.Join(tmp, " ")
}

func main() {
	// fmt.Printf("%q\n", strings.Split(" hello   world!!  ", " "))
	fmt.Printf("%v\n", reverseWords("the sky is blue"))
	fmt.Printf("%v\n", reverseWords("  hello world!  "))
	fmt.Printf("%v\n", reverseWords("a good   example"))
}
