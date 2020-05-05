package main

// https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/
func cuttingRope(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	}
	m := n / 3
	leaf := n % 3
	switch leaf {
	case 0:
		return pow(3, m) % (1e9 + 7)
	case 1:
		return pow(3, m-leaf) * pow(4, leaf) % (1e9 + 7)
	default:
		return pow(3, m) * pow(2, 3-leaf) % (1e9 + 7)
	}
}

func pow(a, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
		ret %= 1e9 + 7
	}
	return ret
}
