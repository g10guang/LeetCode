package main

// https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
func printNumbers(n int) []int {
	end := pow(10, n)
	ret := make([]int, end-1)
	for i := 1; i < end; i++ {
		ret[i-1] = i
	}
	return ret
}

func pow(a, b int) int {
	var ret = 1
	for ; b > 0; b-- {
		ret *= a
	}
	return ret
}
