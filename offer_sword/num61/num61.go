package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
func isStraight(nums []int) bool {
	length := len(nums)
	if length != 5 {
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	jokerCnt := 0
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			jokerCnt++
		} else {
			break
		}
	}
	diff := 0
	for j := jokerCnt + 1; j < length; j++ {
		if nums[j] == nums[j-1] {
			// cannot be equal
			return false
		}
		diff += nums[j] - nums[j-1]
	}
	fmt.Printf("diff=%d jokerCnt=%d\n", diff, jokerCnt)
	return diff <= 4+jokerCnt
}

func main() {
	fmt.Printf("%v\n", isStraight([]int{1, 2, 0, 4, 5}))
	fmt.Printf("%v\n", isStraight([]int{0, 0, 1, 2, 5}))
}
