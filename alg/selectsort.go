package alg

// 选择排序
func SelectSort(elems []int) {
	for i := 0; i < len(elems)-1; i++ {
		min := i
		for j := i+1; j < len(elems); j++ {
			if elems[j] < elems[min] {
				min = j
			}
		}
		elems[i], elems[min] = elems[min], elems[i]
	}
}
