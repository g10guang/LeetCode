package main

// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
func spiralOrder(matrix [][]int) []int {
	var ret []int
	rowLen := len(matrix)
	if rowLen == 0 {
		return ret
	}
	colLen := len(matrix[0])
	acc, total := 0, rowLen*colLen
	minRow, maxRow, minCol, maxCol := 0, rowLen-1, 0, colLen-1
	i, j := 0, 0
	for i >= minRow && i <= maxRow && j >= minCol && j <= maxCol {
		for ; j <= maxCol; j++ {
			ret = append(ret, matrix[i][j])
			acc++
		}
		if acc == total {
			break
		}
		j--
		i++
		minRow++
		for ; i <= maxRow; i++ {
			ret = append(ret, matrix[i][j])
			acc++
		}
		if acc == total {
			break
		}
		i--
		j--
		maxCol--
		for ; j >= minCol; j-- {
			ret = append(ret, matrix[i][j])
			acc++
		}
		if acc == total {
			break
		}
		j++
		i--
		maxRow--
		for ; i >= minRow; i-- {
			ret = append(ret, matrix[i][j])
			acc++
		}
		if acc == total {
			break
		}
		i++
		j++
		minCol++
	}
	return ret
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	spiralOrder(matrix)
}
