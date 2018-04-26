package backtracking

// https://leetcode.com/problems/permutations/description/
func permute(nums []int) [][]int {
	length := len(nums)
	// 记录哪些位置已经被确定
	indexSet := make([]int, length)
	result := make([][]int, countPermutationLength(length))
	resultIndex := 0
	var handle func(index int, position int)
	handle = func(index int, position int) {
		if position >= length {
			return
		}
		if index == length {
			ts := make([]int, length)
			for j := 0; j < length; j++ {
				ts[j] = nums[indexSet[j]]
			}
			result[resultIndex] = ts
			resultIndex++
			return
		}
		if indexSet[position] == -1 {
			indexSet[position] = index
			// 从下一个位置的 0 位置开始
			handle(index+1, 0)
			indexSet[position] = -1
		}
		if index < length {
			handle(index, position+1)
		}
	}
	clear(indexSet)
	handle(0, 0)
	return result
}

func clear(indexSet []int) {
	for i := 0; i < len(indexSet); i++ {
		indexSet[i] = -1
	}
}

func countPermutationLength(length int) int {
	result := 1
	for i := 2; i <= length; i++ {
		result *= i
	}
	return result
}

var Permute = permute
