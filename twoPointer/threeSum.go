package twoPointer

import "sort"

func ThreeSum(arr []int, sum int) (cnt int) {
	sort.Ints(arr)
	for i := 0; i < len(arr) - 2; i++ {
		l := i + 1
		h := len(arr) - 1
		target := sum - arr[i]
		for l < h {
			s := arr[l] + arr[h]
			if s < target {
				l++
				for l < h && arr[l] == arr[l-1] {
					l++
				}
			} else if s > target {
				h--
				for l < h && arr[h] == arr[h+1] {
					h--
				}
			} else {
				// ==
				cnt++
				l++
				for l < h && arr[l] == arr[l-1] {
					l++
					cnt++
				}
				h--
				for l < h && arr[h] == arr[h+1] {
					h++
					cnt++
				}
			}
		}
	}
	return
}
