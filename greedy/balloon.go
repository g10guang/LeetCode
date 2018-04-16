package greedy

/*
1. 根据每个气球的 end 进行升序排序
2. 每次选出 end 最小的气球，也就是该箭射在 end 位置，扫描后面其他还没爆的气球，start <= end 的都会在这次爆裂
*/
func findMinArrowShots(data [][]int) int {
	if len(data) == 0 {
		return 0
	}
	qsort(data)
	cnt := 0
	curPos := 1
	for i := 0; i < len(data); i++ {
		if data[i][0] <= data[curPos][1] {
			continue
		}
		curPos = i
		cnt++
	}
	return cnt
}

func qsort(data [][]int) {
	if len(data) < 2 {
		return
	}
	elem := data[0]
	head := 0
	for i, tail := 1, len(data)-1; i <= tail; {
		if data[i][1] > elem[1] {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			i++
			head++
		}
	}
	data[head] = elem
	qsort(data[0:head])
	qsort(data[head+1:])
}

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
func ShotBalloon(data [][]int) int {
	data = ChangePos(data)
	data = QuickSort(data)
	count := 0
	for i, c := 0, 0; i < len(data); i += c {
		c = FindMaxCountOneShot(data[i:])
		count++
	}
	return count
}

func ChangePos(data [][]int) [][]int {
	for i := len(data) - 1; i >= 0; i-- {
		if data[i][0] > data[i][1] {
			data[i][0], data[i][1] = data[i][1], data[i][0]
		}
	}
	return data
}

// Check how many ballon can be shot in one shot
// And data[0] must be shot
func FindMaxCountOneShot(data [][]int) int {
	if len(data) <= 0 {
		return 0
	}
	end := data[0][1]
	count := 1
	for i := 1; i < len(data); i++ {
		if data[i][0] < end {
			if data[i][1] < end {
				end = data[i][1]
			}
			count++
		} else {
			break
		}
	}
	return count
}

func QuickSort(data [][]int) [][]int {
	if len(data) < 2 {
		return data
	}
	head := 0
	midElem := data[0]
	mid := midElem[0]
	for i, tail := 1, len(data)-1; i <= tail; {
		if data[i][0] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			i++
			head++
		}
	}
	data[head] = midElem
	QuickSort(data[:head])
	QuickSort(data[head+1:])
	return data
}
