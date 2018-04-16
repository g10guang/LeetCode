package alg

// 归并排序
func MergeSort(elems []int) {
	length := len(elems)
	if length <= 1 {
		return
	}
	mid := length / 2
	MergeSort(elems[:mid])
	MergeSort(elems[mid:])
	// 需要牺牲额外的空间
	ts := make([]int, length)
	copy(ts, elems)
	i, j, head := 0, mid, 0
	for i < mid && j < length {
		if ts[i] <= ts[j] {
			elems[head] = ts[i]
			i++
		} else {
			elems[head] = ts[j]
			j++
		}
		head++
	}
	for i < mid {
		elems[head] = ts[i]
		head++
		i++
	}
	for j < length {
		elems[head] = ts[j]
		head++
		j++
	}
}
