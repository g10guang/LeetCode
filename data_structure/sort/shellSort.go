package sort

// shell sort is a advanced solution for insert sort.
// insert sort switch the adjacent element.
// shell sort can switch faster. and make elem near its final position.
func ShellSort(arr []int) {
	length := len(arr)
	h := 1
	// experience value to make shell sort more effectively.
	for h < length/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		// according to insert sort. arr[0..h] is in order.
		// we just need to sort from arr[h...length-1]
		for i := h; i < length; i++ {
			for j := i; j >= h && arr[j-h] > arr[j]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
}
