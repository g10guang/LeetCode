package alg

// 希尔排序，多次间隔递减至1的插入排序
// 因为插入排序在数据基本有序的情况下效率高
func ShellSort(elems []int) {
	length := len(elems)
	if length <= 1 {
		return
	}
	h := 1
	for h * 3 < length {
		h *= 3
	}
	for h >= 1 {
		//for i := 0; i < h; i++ {
		//	for j := i + h; j < length && elems[j] < elems[j-h]; j++ {
		//		elems[j], elems[j-h] = elems[j-h], elems[j]
		//	}
		//}
		for i := h; i < length; i++ {
			for j := i; j >= h && elems[j] < elems[j-h]; j-- {
				elems[j], elems[j-h] = elems[j-h], elems[j]
			}
		}
		h /= 3
	}
}
