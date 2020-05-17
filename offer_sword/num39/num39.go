package main

// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// func majorityElement(nums []int) int {
// 	mark := make(map[int]int)
// 	for _, n := range nums {
// 		mark[n]++
// 	}
// 	marjority := len(nums)/2 + 1
// 	for num, cnt := range mark {
// 		if cnt >= marjority {
// 			return num
// 		}
// 	}
// 	return -1
// }

func majorityElement(nums []int) int {
	num := 0
	cnt := 0
	for _, n := range nums {
		if cnt == 0 {
			num = n
		}
		if n == num {
			cnt++
		} else {
			cnt--
		}
	}
	return num
}
