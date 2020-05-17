package main

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
func singleNumbers(nums []int) (result []int) {
	mark := make(map[int]int)
	for _, v := range nums {
		mark[v]++
	}
	for k, v := range mark {
		if v != 2 {
			result = append(result, k)
		}
		if len(result) == 2 {
			return result
		}
	}
	return 
}