package alg

func QuickSort(elems []int) {
	if len(elems) <= 1 {
		return
	}
	mid := elems[0]
	head := 0
	for i, tail := 1, len(elems) - 1; i <= tail; {
		if elems[i] > mid {
			elems[i], elems[tail] = elems[tail], elems[i]
			tail--
		} else {
			elems[i], elems[head] = elems[head], elems[i]
			head++
			i++
		}
	}
	elems[head] = mid
	QuickSort(elems[:head])
	QuickSort(elems[head+1:])
}

func QuickSort2(elems []int) {
	if len(elems) <= 1 {
		return
	}
	mid := elems[0]
	head, tail := 1, len(elems)-1
	// 不能使用 for head < tail，不然两个元素的时候有可能处触发 big small 的顺序
	for true {
		for elems[head] <= mid && head < tail {
			head++
		}
		for elems[tail] >= mid && tail > 0 {
			tail--
		}
		if head < tail {
			elems[head], elems[tail] = elems[tail], elems[head]
		} else {
			break
		}
	}
	elems[tail], elems[0] = elems[0], elems[tail]
	QuickSort2(elems[:tail])
	QuickSort2(elems[tail+1:])
}

// 使用3向切分方法进行快速排序，适用于有大量重复元素的排序
func Quick3WaySort(elems []int) {
	if len(elems) <= 1 {
		return
	}
	mid := elems[0]
	i, lt, gt := 1, 0, len(elems)-1
	for i <= gt {
		if elems[i] > mid {
			elems[gt], elems[i] = elems[i], elems[gt]
			gt--
		} else if elems[i] < mid {
			elems[lt], elems[i] = elems[i], elems[lt]
			lt++
			i++
		} else {
			// elems[i] == mid
			i++
		}
	}
	Quick3WaySort(elems[:lt])
	Quick3WaySort(elems[gt+1:])
}
