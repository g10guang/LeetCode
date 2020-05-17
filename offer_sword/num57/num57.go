package main

// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
func twoSum2(nums []int, target int) []int {
	mark := make(map[int]int)
	for _, v := range nums {
		mark[v]++
	}
	for _, v := range nums {
		if v + v == target {
			if mark[v] > 1 {
				return []int{v, v}
			}
		}
		if mark[target-v] > 0 {
			return []int{v, target - v}
		}
	}
	return nil
}


func twoSum(nums []int, target int) []int {
	low, high := 0, len(nums) - 1
	for low < high {
		sum := nums[low] + nums[high]
		if sum == target {
			return []int{nums[low], nums[high]}
		} else if sum > target {
			high--
		} else {
			low++
		}
	}
	return nil
}