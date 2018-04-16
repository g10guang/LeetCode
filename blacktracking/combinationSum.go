package blacktracking

var CombinationSum = combinationSum

// https://leetcode.com/problems/combination-sum/description/
func combinationSum(candidates []int, target int) [][]int {
	// 因为这里确保 candidates element > 0
	var handle func(int, int)
	result := make([][]int, 0, 16)
	tmpResult := make([]int, 0, 4)
	// index 用于控制前面的元素先被访问，也就是去除相同集和之间的排序
	handle = func(index int, sum int) {
		if sum == target {
			t := make([]int, len(tmpResult))
			copy(t, tmpResult)
			result = append(result, t)
		}
		if sum > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			if sum+candidates[i] <= target {
				tmpResult = append(tmpResult, candidates[i])
				handle(i, sum + candidates[i])
				tmpResult = tmpResult[:len(tmpResult)-1]
			}
		}
	}
	handle(0, 0)
	return result
}
