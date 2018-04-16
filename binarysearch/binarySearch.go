package binarysearch

func binarySearch(data []int, target int) int {
	head, tail := 0, len(data) - 1
	for head <= tail {
		mid := head + (tail - head) / 2
		if data[mid] < target {
			tail = mid - 1
		} else if data[mid] > target {
			head = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
