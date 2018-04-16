package graph

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row <= 0 {
		return 0
	}
	col := len(grid[0])
	mark := make([][]bool, row)
	for i := 0; i < row; i++ {
		mark[i] = make([]bool, col)
	}
	islandNum := 0
	var handle func(x, y int) bool
	handle = func(x, y int) bool {
		if x < 0 || x >= row || y < 0 || y >= col {
			return false
		}
		if grid[x][y] == 0 {
			return false
		}
		if mark[x][y] {
			return false
		}
		mark[x][y] = true
		handle(x-1, y)
		handle(x+1, y)
		handle(x, y-1)
		handle(x, y+1)
		return true
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if handle(i, j) {
				islandNum++
			}
		}
	}
	return islandNum
}

var NumIslands = numIslands
