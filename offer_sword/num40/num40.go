package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers1(arr []int, k int) []int {
	result := make([]int, k)
	if k <= 0 {
		return result
	}
	copy(result, arr)
	for _, num := range arr[k:] {
		maxIdx := 0
		for idx, v := range result {
			if v > result[maxIdx] {
				maxIdx = idx
			}
		}
		if result[maxIdx] > num {
			result[maxIdx] = num
		}
	}
	return result
}

// qucik sort
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == k {
		return arr
	}
	if k == 0 {
		return []int{}
	}
	var partition func(start, end int) []int
	partition = func(start, end int) []int {
		pivot := arr[end]
		index := start
		for j := start; j < end; j++ {
			if arr[j] < pivot {
				arr[index], arr[j] = arr[j], arr[index]
				index++
			}
		}
		arr[index], arr[end] = arr[end], arr[index]
		if index == k-1 {
			return arr[:k]
		} else if index > k-1 {
			return partition(start, index-1)
		} else {
			return partition(index+1, end)
		}
	}
	return partition(0, len(arr)-1)
}

func main() {
	// arr := []int{0, 1, 2, 1}
	arr := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	result := getLeastNumbers(arr, 8)
	fmt.Printf("result=%+v\n", result)
}
