package twoPointer

import "sort"

// give an int slice, find how many 2-size pair which sum is 0
// save storage
func ZeroSum(arr []int, sum int) (cnt int) {
	sort.Ints(arr) // use go build-in sort function to sort.
	p, q := 0, len(arr)-1
	for p < q {
		s := arr[p] + arr[q]
		if s == sum {
			cnt++
			p++ // make sure a step move in every loop.
		} else if s > sum {
			q--
		} else {
			p++
		}
	}
	return
}

// save time
func ZeroSum2(arr []int, sum int) (cnt int) {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		cnt += v * m[sum-k]
	}
	// step above calculate 2 * cnt
	cnt /= 2
	return
}
