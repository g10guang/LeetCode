package backtracking

var Combine = combine

// https://leetcode.com/problems/combinations/description/
func combine(n, k int) [][]int {
	tmpResult := make([]int, 0, k)
	result := make([][]int, 0, countLength(n, k))
	var handle func(int)
	handle = func(start int) {
		// 控制递归的深度
		if len(tmpResult) == k {
			t := make([]int, k)
			copy(t, tmpResult)
			result = append(result, t)
		}
		for i := start; i <= n; i++ {
			tmpResult = append(tmpResult, i)
			handle(i+1)
			// 确保 i 就在栈顶
			tmpResult = tmpResult[:len(tmpResult)-1]
		}
	}
	handle(1)
	return result
}

// 计算组合数
func countLength(n, k int) int {
	x, y := 1, 1
	for k > 0 {
		x *= n
		n--
		y *= k
		k--
	}
	return x / y
}
