package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
func maxProfit(prices []int) int {
	cost, profit := math.MaxInt64, 0
	for _, p := range prices {
		cost = min(cost, p)
		profit = max(profit, p - cost)
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%v\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v\n", maxProfit([]int{7, 6, 4, 3, 1}))
}
