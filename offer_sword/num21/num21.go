package main

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func exchange(nums []int) []int {
	for low, high := 0, len(nums)-1; low < high; low, high = low+1, high-1 {
		for nums[low]%2 == 1 && low < high {
			low++
		}
		for nums[high]%2 == 0 && low < high {
			high--
		}
		if low >= high {
			break
		}
		nums[low], nums[high] = nums[high], nums[low]
	}
	return nums
}
