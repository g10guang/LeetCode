package greedy

// https://leetcode.com/problems/non-decreasing-array/description/
func CheckPossibility2(nums []int) bool {
	length := len(nums)
	data := make([]int, length)
	copy(data, nums)
	hasModified := false
	for i := 1; i < length; i++ {
		if data[i] < data[i-1] {
			if hasModified {
				return false
			}
			hasModified = true
			// 两种情况：
			// 1. nums[i-1] 变小
			// 2. nums[i] 变大
			if i < length-1 && data[i-1] > data[i+1] {
				if i >= 2 {
					if data[i-2] > data[i] {
						return false
					}
					data[i-1] = data[i-2]
				} else {
					data[i-1] = data[i]
				}
			} else {
				data[i] = data[i-1]
			}
		}
	}
	return true
}

func CheckPossibility(nums []int) bool {
	hasModified := false
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < last {
			if hasModified {
				return false
			}
			hasModified = true
			i++
			if i >= len(nums) {
				break
			}
			if i-3 >= 0 {
				if nums[i] < nums[i-3] {
					return false
				}
			}
		}
		last = nums[i]
	}
	return true
}
