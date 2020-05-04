package main

import "fmt"

// Q: https://leetcode-cn.com/problems/edit-distance/

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1), len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2), len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var distant func(i, j int) int

	distant = func(i, j int) int {
		if i == -1 {
			return j + 1
		} else if j == -1 {
			return i + 1
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if word1[i] == word2[j] {
			v := min(distant(i-1, j-1), distant(i-1, j)+1, distant(i, j-1)+1)
			memo[i][j] = v
			return v
		}

		v := min(distant(i-1, j)+1, distant(i, j-1)+1, distant(i-1, j-1)+1)
		memo[i][j] = v
		return v
	}
	return distant(len(word1)-1, len(word2)-1)
}

func min(n int, nums ...int) int {
	v := n
	for _, x := range nums {
		if v > x {
			v = x
		}
	}
	return v
}

func main() {
	fmt.Printf("result1=%v\n", minDistance("hello", "helo"))
	fmt.Printf("result1=%v\n", minDistance("", "helo"))
	fmt.Printf("result1=%v\n", minDistance("", ""))
	fmt.Printf("result1=%v\n", minDistance("hello", "hello"))
}
