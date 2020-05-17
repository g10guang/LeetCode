package main

import "fmt"

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
func missingNumber(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		v := nums[i]
		if v >= length {
			continue
		}
		if i == v {
			continue
		}
		nums[i] = nums[v]
		nums[v] = v
		i--
	}
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return length
}

func main() {
	fmt.Printf("%v\n", missingNumber([]int{1, 2, 3}))
	fmt.Printf("%v\n", missingNumber([]int{1, 2, 0}))
}