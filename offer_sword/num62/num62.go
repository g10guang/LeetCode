package main

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/solution/javajie-jue-yue-se-fu-huan-wen-ti-gao-su-ni-wei-sh/
// Copy others answer
func lastRemaining(n int, m int) int {
	ans := 0
	for i := 2; i < n; i++ {
		ans = (ans + m) % i
	}
	return ans
}