package main

// https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
func myPow(x float64, n int) float64 {
	var ret = 1.0
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n > 0 {
		if n&1 == 1 {
			ret *= x
		}
		x *= x
		n >>= 1
	}
	return ret
}
