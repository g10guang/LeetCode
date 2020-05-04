package main

import (
	"fmt"
	"sort"
)

// Q: https://leetcode-cn.com/problems/circus-tower-lcci/

func bestSeqAtIndex(height []int, weight []int) int {
	if len(height) == 0 || len(height) != len(weight) {
		return 0
	}
	slice := &MySlice{
		height: height,
		weight: weight,
	}
	sort.Sort(slice)
	memo := make([]int, slice.Len(), slice.Len())
	memo[0] = 1
	for i := 1; i < slice.Len(); i++ {
		memo[i] = 1
		for j := 0; j < i; j++ {
			if slice.Less(j, i) {
				memo[i] = max(memo[i], memo[j]+1)
			}
		}
	}
	return max(0, memo...)
}

func max(n int, nums ...int) int {
	ret := n
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

type MySlice struct {
	height []int
	weight []int
}

func (s *MySlice) Less(i, j int) bool {
	return s.height[i] < s.height[j] && s.weight[i] < s.weight[j]
}

func (s *MySlice) Swap(i, j int) {
	s.height[i], s.height[j] = s.height[j], s.height[i]
	s.weight[i], s.weight[j] = s.weight[j], s.weight[i]
}

func (s *MySlice) Len() int {
	return len(s.height)
}

func main() {
	fmt.Printf("result1=%v\n", bestSeqAtIndex([]int{65, 70, 56, 75, 60, 68}, []int{100, 150, 90, 190, 95, 110}))
	fmt.Printf("result2=%v\n", bestSeqAtIndex([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Printf("result3=%v\n", bestSeqAtIndex([]int{2868, 5485, 1356, 1306, 6017, 8941, 7535, 4941, 6331, 6181}, []int{5042, 3995, 7985, 1651, 5991, 7036, 9391, 428, 7561, 8594}))
	fmt.Printf("result4=%v\n", bestSeqAtIndex([]int{5401, 9628, 3367, 6600, 6983, 7853, 5715, 2654, 4453, 8619}, []int{3614, 1553, 2731, 7894, 8689, 182, 7632, 4465, 8932, 4304}))
}
