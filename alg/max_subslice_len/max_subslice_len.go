package main

func shortestSubarray(A []int, K int) int {
	acc := 0
	maxLength := 0
	for i, v := range A {
		if acc >= 0 && acc+v >= 0 {
			maxLength++
		} else if acc < 0 {

		}
	}
	if maxLength == 0 || acc < K {
		return -1
	}
	return maxLength
}
