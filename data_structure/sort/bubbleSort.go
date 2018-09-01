package sort

func BubbleSort(arr []int) {
	// move the biggest number to the end.
	isOrder := false
	for i := len(arr); i > 0 && !isOrder; i-- {
		isOrder = true
		for j := 1; j < i; j++ {
			// if in one turn no elem switch. it is order for all elem.
			// so if the data is almost in order, bubble sort is fast.
			if arr[j-1] > arr[j] {
				if isOrder {
					isOrder = false
				}
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
