package blacktracking

var PermuteUnique = permuteUnique2

// https://leetcode.com/problems/permutations-ii/description/
func permuteUnique(nums []int) [][]int {
	length := len(nums)
	// 记录哪些位置已经被确定
	indexSet := make([]int, length)
	result := make([][]int, 0, countPermutationLength(length))
	// 记录每个元素正在被访问的位置
	curPos := make([]int, length)
	clear(curPos)
	indexMap := scanGenerateMap(nums)
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
			result = append(result, ts)
			return
		}
		sameIndex := indexMap[nums[index]]
		for _, v := range sameIndex {
			if v >= index {
				break
			}
			if curPos[v] >= position {
				handle(index, curPos[v]+1)
				return
			}
		}
		if indexSet[position] == -1 {
			indexSet[position] = index
			curPos[index] = position
			// 从下一个位置的 0 位置开始
			handle(index+1, 0)
			indexSet[position] = -1
			curPos[index] = -1
		}
		if index < length {
			handle(index, position+1)
		}
	}
	clear(indexSet)
	handle(0, 0)
	return result
}

// 扫描，然后生成每个字符出现的所有位置
func scanGenerateMap(nums []int) map[int][]int {
	indexMap := make(map[int][]int)
	for i, v := range nums {
		_, ok := indexMap[v]
		if !ok {
			indexMap[v] = []int{i}
		} else {
			indexMap[v] = append(indexMap[v], i)
		}
	}
	return indexMap
}

/*
这个算法的核心思想：
1. 先对元素进行排序，经过排序之后，相同的元素会相邻
2. 相同元素必须保证前一个元素先于后面元素被访问，这一点消除了相同元素之前的排列问题
*/
func permuteUnique2(nums []int) [][]int {
	length := len(nums)
	result := make([][]int, 0)
	qsort(nums)
	tmpResult := make([]int, 0, length)
	visited := make([]bool, length)
	var handle func()
	handle = func() {
		// 完成了一次的递归，这个条件用于控制递归栈的深度
		if len(tmpResult) == length {
			t := make([]int, length)
			copy(t, tmpResult)
			result = append(result, t)
		}
		for i := 0; i < length; i++ {
			// i 位置的元素已经加入到了本次递归中
			if visited[i] {
				continue
			}
			// 必须保证相同的元素前一个元素先于后一个元素加入到本地递归
			if i > 0 && nums[i-1] == nums[i] && !visited[i-1] {
				continue
			}
			visited[i] = true
			tmpResult = append(tmpResult, nums[i])
			handle()
			// 从递归栈中消除当前元素的状态
			tmpResult = tmpResult[:len(tmpResult)-1]
			visited[i] = false
		}
	}
	handle()
	return result
}

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	midElem := nums[0]
	head := 0
	for i, tail := 1, len(nums) - 1; i <= tail; {
		if nums[i] > midElem {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			i++
			head++
		}
	}
	nums[head] = midElem
	qsort(nums[:head])
	qsort(nums[head+1:])
}
