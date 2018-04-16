package graph

type Position struct {
	x int
	y int
}


func maxAreaOfIsland(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	// mark which point has visited or not.
	mark := make([][]bool, row)
	for i := 0; i < row; i++ {
		mark[i] = make([]bool, col)
		for j := 0; j < col; j++ {
			mark[i][j] = false
		}
	}
	max := 0
	cur := 0
	var dfsRecursive func(x, y int) int
	dfsRecursive = func(x, y int) int {
		if x >= row || x < 0 || y >= col || y < 0 {
			return 0
		}
		if mark[x][y] {
			return 0
		}
		mark[x][y] = true
		if grid[x][y] == 0 {
			return 0
		}
		return 1 + dfsRecursive(x-1, y) + dfsRecursive(x+1, y) + dfsRecursive(x, y-1) + dfsRecursive(x, y+1)
	}
	// visit every point of grid
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cur = dfsRecursive(i, j)
			if cur > max {
				max = cur
			}
		}
	}
	return max
}

var MaxAreaOfIsland = maxAreaOfIsland

func dfs(grid [][]int, mark [][]bool, i, j, row, col int) ([][]bool, int) {
	counter := 1
	stack := make([]*Position, 1)
	stack[0] = &Position{x:i, y:j}
	handle := func(x, y int) {
		// valid the position
		if x >= 0 && x < row && y >= 0 && y < col {
			if grid[x][y] == 1 && !mark[x][y] {
				counter++
				mark[x][y] = true
				stack = append(stack, &Position{x:x, y:y})
			}
		}
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left, right, up, down := node.y - 1, node.y + 1, node.x - 1, node.x + 1
		handle(node.x, left)
		handle(node.x, right)
		handle(up, node.y)
		handle(down, node.y)
	}
	return mark, counter
}
