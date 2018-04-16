package twoPointer

// https://leetcode.com/problems/merge-sorted-array/description/
func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	if len(nums1) < m+n {
		nums1 = append(nums1, nums2...)
	}
	k := m + n - 1
	i, j := m-1, n-1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for i >= 0 && i != k {
		nums1[k] = nums1[i]
		k--
		i--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
	return nums1
}


