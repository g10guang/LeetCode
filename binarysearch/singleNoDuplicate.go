package binarysearch

// https://leetcode.com/problems/single-element-in-a-sorted-array/description/
func singleNonDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	l, h := 0, len(nums) - 1
	// if l == h then nums[m + 1] may be out of range.
	for l < h {
		m := l + (h - l) / 2
		// keep m to be even
		if m % 2 == 1{
			m--
		}
		if nums[m] == nums[m+1] {
			// single one should be on the right of m
			l = m + 2
		} else {
			// single one should be on the left of m
			h = m
		}
	}
	return nums[l]
}

var SingleNonDuplicate = singleNonDuplicate
