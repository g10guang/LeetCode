package alg

// 插入排序
func InsertSort(elems []int) {
	for i := 1; i < len(elems); i++ {
		for j := i; j > 0; j-- {
			if elems[j] < elems[j-1] {
				elems[j], elems[j-1] = elems[j-1], elems[j]
			} else {
				break
			}
		}
	}
}