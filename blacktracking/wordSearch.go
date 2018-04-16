package blacktracking

// https://leetcode.com/problems/word-search/description/
func exist(board [][]byte, word string) bool {
	row := len(board)
	if row <= 0 {
		return false
	}
	col := len(board[0])
	mark := make([][]bool, row)
	for i := 0; i < row; i++ {
		mark[i] = make([]bool, col)
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var handle func(x, y, start int) bool
	handle = func(x, y, start int) bool {
		if x < 0 || x >= row || y < 0 || y >= col {
			return false
		}
		if mark[x][y] {
			return false
		}
		if board[x][y] != word[start] {
			return false
		}
		if start == len(word) - 1 {
			return true
		}
		mark[x][y] = true
		flag := false
		for _, v := range directions {
			if handle(x+v[0], y+v[1], start+1) {
				flag = true
				break
			}
		}
		mark[x][y] = false
		return flag
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if handle(i, j, 0) {
				return true
			}
		}
	}
	return false
}

var Exist = exist
