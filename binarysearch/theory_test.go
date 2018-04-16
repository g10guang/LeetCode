package binarysearch_test

import (
	"testing"
)

func Search(nums []int, target int) int {
	l, h := 0, len(nums)
	for l <= h {
		mid := l + (h-l)/2
		midValue := nums[mid]
		if midValue > target {
			h = mid - 1
		} else if midValue < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	result := Search(input, 5)
	// t.Println(result)
	if result != 4 {
		t.Errorf("result=%d should be %d", result, 4)
	}
}
