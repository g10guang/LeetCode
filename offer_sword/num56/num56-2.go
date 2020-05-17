package main

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
func singleNumber(nums []int) int {
	mark := make(map[int]int)
	for _, v := range nums {
		mark[v]++
	}
	for k, v := range mark {
		if v != 3 {
			return k
		}
	}
	return 0
}