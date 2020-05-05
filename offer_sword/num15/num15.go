package main

// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	ret := 0
	for num > 0 {
		ret += num % 2
		num /= 2
	}
	return ret
}