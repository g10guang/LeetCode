package main

// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
var memo [101]int

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	val := memo[n]
	if val > 0 {
		return val
	}
	val = fib(n-2) + fib(n-1)
	val %= 1e9 + 7
	memo[n] = val
	return val
}
