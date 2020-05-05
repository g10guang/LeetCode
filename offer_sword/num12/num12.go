package main

// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
func exist(board [][]byte, word string) bool {
	strLen := len(word)
	rowNum, colNum := len(board), len(board[0])
	memo := make([][]bool, len(board))
	for i := range memo {
		memo[i] = make([]bool, colNum)
	}
	var fn func(i, j, k int) bool
	fn = func(i, j, k int) bool {
		if k == strLen {
			return true
		}
		if i < 0 || i >= rowNum || j < 0 || j >= colNum {
			return false
		}
		if memo[i][j] {
			return false
		}
		if board[i][j] != word[k] {
			return false
		}
		memo[i][j] = true
		if fn(i-1, j, k+1) || fn(i+1, j, k+1) || fn(i, j-1, k+1) || fn(i, j+1, k+1) {
			return true
		}
		memo[i][j] = false
		return false
	}
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			if fn(row, col, 0) {
				return true
			}
		}
	}
	return false
}
