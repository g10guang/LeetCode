package main

// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
func movingCount(m int, n int, k int) int {
	memo := make([][]bool, m)
	for i := range memo {
		memo[i] = make([]bool, n)
	}
	var cnt int
	var fn func(i, j int)
	fn = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if memo[i][j] {
			return
		}
		if sum(i, j) > k {
			return
		}
		cnt++
		memo[i][j] = true
		fn(i-1, j)
		fn(i+1, j)
		fn(i, j-1)
		fn(i, j+1)
	}
	fn(0, 0)
	return cnt
}

func sum(i, j int) int {
	ret := 0
	for i > 0 {
		ret += i % 10
		i /= 10
	}
	for j > 0 {
		ret += j % 10
		j /= 10
	}
	return ret
}
