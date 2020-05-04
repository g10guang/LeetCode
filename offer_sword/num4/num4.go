package main

// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		val := matrix[row][col]
		switch {
		case val == target:
			return true
		case val > target:
			col--
		case val < target:
			row++
		}
	}
	return false
}
