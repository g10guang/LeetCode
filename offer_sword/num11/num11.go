package main

// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
func minArray(numbers []int) int {
	return pivot(numbers)
}

func pivot(nums []int) int {
	length := len(nums)
	switch length {
	case 1, 2, 3:
		return min(nums...)
	}
	low, high := 0, length-1
	mid := (low + high) / 2
	lowVal, midVal := nums[low], nums[mid]
	highVal := nums[high]
	switch {
	case lowVal > midVal || midVal <= highVal:
		return pivot(nums[:mid+1])
	}
	return min(pivot(nums[:mid+1]), pivot(nums[mid+1:]))
}

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("empty nums")
	}
	ret := nums[0]
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}
