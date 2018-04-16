package graph

// https://leetcode.com/problems/pacific-atlantic-water-flow/description/

var PacificAtlantic = pacificAtlantic

func pacificAtlantic(matrix [][]int) [][]int {
	result := make([][]int, 0)
	row := len(matrix)
	if row <= 0 {
		return result
	}
	col := len(matrix[0])
	mark := make([][]byte, row)
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < row; i++ {
		mark[i] = make([]byte, col)
	}
	for i := 0; i < row; i++ {
		mark[i][0] |= 1		// 可达太平洋
		mark[i][col-1] |= 2	// 可达大西洋
	}
	for j := 0; j < col; j++ {
		mark[0][j] |= 1		// 可达太平洋
		mark[row-1][j] |= 2	// 可达大西洋
	}
	var dfs func(x, y, lastHeight int, area byte)
	dfs = func(x, y, lastHeight int, area byte) {
		if x < 0 || x >= row || y < 0 || y >= col {
			return
		}
		if matrix[x][y] < lastHeight || mark[x][y] & area != 0 {
			return
		}
		mark[x][y] |= area	// 该点可达 area 区域
		for _, v := range directions {
			dfs(x+v[0], y+v[1], matrix[x][y], area)
		}
	}
	var start func(x, y int, area byte)
	start = func(x, y int, area byte) {
		for _, v := range directions {
			dfs(x+v[0], y+v[1], matrix[x][y], area)
		}
	}
	for i := 0; i < row; i++ {
		start(i, 0, 1)
		start(i, col-1, 2)
	}
	for j := 0; j < col; j++ {
		start(0, j, 1)
		start(row-1, j, 2)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mark[i][j] & 3 == 3 {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}
