package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func maxSubArray(nums []int) int {
	maxVal := int(math.MinInt64)
	curSum := 0
	for _, v := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += v
		if curSum > maxVal {
			maxVal = curSum
		}
	}
	return maxVal
}

func main() {
	fmt.Printf("%d\n", maxSubArray([]int{}))
	fmt.Printf("%d\n", maxSubArray([]int{-2, -1}))
	fmt.Printf("%d\n", maxSubArray([]int{1, 2, 3, 4}))
	fmt.Printf("%d\n", maxSubArray([]int{1, 2, 3, -4}))
}
