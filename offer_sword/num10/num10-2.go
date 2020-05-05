package main

var memo [101]int

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func numWays(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	}
	val := memo[n]
	if val > 0 {
		return val
	}
	val = numWays(n-1) + numWays(n-2)
	val %= 1e9 + 7
	memo[n] = val
	return val
}
