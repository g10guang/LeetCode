package twoPointer

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum(numbers []int, target int) []int {
	p, q := 0, len(numbers) - 1
	for p < q && numbers[p] + numbers[q] != target {
		if numbers[p] + numbers[q] > target {
			q--
		} else {
			p++
		}
	}
	return []int{p+1, q+1}
}
