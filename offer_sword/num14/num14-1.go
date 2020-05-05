package main

// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
func cuttingRope(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	}
	m := n / 3
	leaf := n % 3
	if leaf == 0 {
		return pow(3, m)
	}
	switch leaf {
	case 0:
		return pow(3, m)
	case 1:
		return pow(3, m-leaf) * pow(4, leaf)
	default:
		return pow(3, m) * pow(2, 3-leaf)
	}
}

func pow(a, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}
