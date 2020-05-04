package main

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func findRepeatNumber(nums []int) int {
	for i, v := range nums {
		if v == i {
			continue
		}
		if nums[v] == v {
			return v
		}
		nums[v], nums[i] = v, nums[v]
	}
	return 0
}