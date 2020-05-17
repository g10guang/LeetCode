package main

// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
func maxValue(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}
	mark := make([][]int, row)
	for i := range mark {
		mark[i] = make([]int, col)
	}
	mark[0][0] = grid[0][0]
	max := func(i, j int) int {
		if i == 0 && j == 0 {
			return grid[i][j]
		}
		if i == 0 {
			return mark[i][j-1] + grid[i][j]
		}
		if j == 0 {
			return mark[i-1][j] + grid[i][j]
		}
		return maxFn(mark[i][j-1], mark[i-1][j]) + grid[i][j]
	}
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			val := max(x, y)
			if mark[x][y] < val {
				mark[x][y] = val
			}
		}
	}
	return mark[row-1][col-1]
}

func maxFn(a, b int) int {
	if a > b {
		return a
	}
	return b
}