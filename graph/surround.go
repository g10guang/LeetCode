package graph

// https://leetcode.com/problems/surrounded-regions/description/
func solve(board [][]byte) {
	row := len(board)
	if row <= 0 {
		return
	}
	col := len(board[0])
	// 1-->has visited 0-->not visited 2-->visiting
	mark := make([][]byte, row)
	for i := 0; i < row; i++ {
		mark[i] = make([]byte, col)
	}
	var isSurrounded func(i, j int) bool
	// judge whether is surround by 'X' or not.
	isSurrounded = func(i, j int) bool {
		// use DFS try to find 'O' in boundary
		if i < 0 || i >= row || j < 0 || j >= col {
			return false
		}
		if board[i][j] == 'X' {
			return true
		}
		if mark[i][j] == 1 {
			return false
		} else if mark[i][j] == 2 {
			return true
		}
		mark[i][j] = 2
		r := isSurrounded(i-1, j) && isSurrounded(i+1, j) && isSurrounded(i, j-1) && isSurrounded(i, j+1)
		if r {
			board[i][j] = 'X'
		}
		mark[i][j] = 1
		return r
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' && mark[i][j] != 1 {
				isSurrounded(i, j)
			}
		}
	}
}

var Solve = solve2

// 思路：从最外层往里面层进行深度有限搜索，将所有已经搜索过的打标记
func solve2(board [][]byte) {
	row := len(board)
	if row <= 0 {
		return
	}
	col := len(board[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= row || y < 0 || y >= col {
			return
		}
		if board[x][y] == 'O' {
			board[x][y] = 'T'
			dfs(x-1, y)
			dfs(x+1, y)
			dfs(x, y-1)
			dfs(x, y+1)
		}
	}
	for i := 0; i < row; i++ {
		dfs(i, 0)
		dfs(i, col-1)
	}
	for i := 0; i < col; i++ {
		dfs(0, i)
		dfs(row-1, i)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch board[i][j] {
			case 'T':
				board[i][j] = 'O'
			default:
				board[i][j] = 'X'
			}
		}
	}
}
