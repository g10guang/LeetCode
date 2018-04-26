package backtracking

var CombinationSum3 = combinationSum3

// https://leetcode.com/problems/combination-sum-iii/description/
func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0, 8)
	tmpResult := make([]int, 0, 9)
	var handle func(int, int)
	handle = func(index int, sum int) {
		if sum == n && len(tmpResult) == k {
			t := make([]int, len(tmpResult))
			copy(t, tmpResult)
			result = append(result, t)
		}
		if sum > n || len(tmpResult) >= k {
			return
		}
		for i := index; i <= 9; i++ {
			tmpResult = append(tmpResult, i)
			handle(i+1, sum+i)
			tmpResult = tmpResult[:len(tmpResult)-1]
		}
	}
	handle(1, 0)
	return result
}
