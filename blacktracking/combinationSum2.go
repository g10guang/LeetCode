package blacktracking

var CombinationSum2 = combinationSum2

func combinationSum2(candidates []int, target int) [][]int {
	qsort(candidates)	// 保证相同的元素相邻
	// 因为这里确保 candidates element > 0
	var handle func(int, int)
	result := make([][]int, 0, 16)
	tmpResult := make([]int, 0, 4)
	visited := make([]bool, len(candidates))
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
			// 如果相同元素上一个元素没有 visited 则不访问该元素
			// 相同元素前一个已经被访问过了，本次访问就不需要再继续
			if i > 0 && candidates[i-1] == candidates[i] && !visited[i-1] {
				continue
			}
			if sum+candidates[i] <= target {
				visited[i] = true
				tmpResult = append(tmpResult, candidates[i])
				handle(i+1, sum + candidates[i])
				tmpResult = tmpResult[:len(tmpResult)-1]
				visited[i] = false
			}
		}
	}
	handle(0, 0)
	return result
}
