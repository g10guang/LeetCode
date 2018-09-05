package sort

// sort from arr[0..mid-1] arr[mid..len(arr)-1]
func mergeSort(first []int, second []int) bool {
	length := len(first)
	mid := length / 2
	var flag bool
	if length > 1 {
		mergeSort(first[0:mid], second[0:mid])
		flag = mergeSort(first[mid:], second[mid:])
	}
	if flag {
		// value-copy for slice is costless. because it record the pointer not array.
		first, second = second, first
	}
	// after last recursive call.
	// the result is on second
	i, j, p := 0, mid, 0
	for ; i < mid && j < length; p++ {
		if first[i] <= first[j] {
			second[p] = first[i]
			i++
		} else {
			second[p] = first[j]
			j++
		}
	}
	for ; i < mid; i, p = i+1, p+1 {
		second[p] = first[i]
	}
	for ; j < length; j, p = j+1, p+1 {
		second[p] = first[j]
	}
	//copy(first, second)
	return !flag
}

func MergeSort(arr []int) {
	// use a temperature array to merge.
	t := make([]int, len(arr))
	mergeSort(arr, t)
}
