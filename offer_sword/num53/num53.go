package main

import "fmt"

// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 1
	}
	low, high := 0, len(nums)-1
	var idx = -1
	for low <= high {
		mid := (low + high) / 2
		val := nums[mid]
		if val > target {
			high = mid - 1
		} else if val < target {
			low = mid + 1
		} else {
			idx = mid
			break
		}
	}
	if idx < 0 {
		return 0
	}
	sum := 0
	for i := idx; i >= 0; i-- {
		if nums[i] == target {
			sum++
		} else {
			break
		}
	}
	for j := idx + 1; j < len(nums); j++ {
		if nums[j] == target {
			sum++
		} else {
			break
		}
	}
	return sum
}

func main() {
	fmt.Printf("%v\n", search([]int{2, 2}, 2))
	fmt.Printf("%v\n", search([]int{1, 4}, 4))
}
