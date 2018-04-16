package graph

func findCircleNum(M [][]int) int {
	n := len(M)
	mark := make([]bool, n)
	circle := 0
	var dfs func(a, b int) bool
	dfs = func(a, b int) bool {
		if a < 0 || a >= n || b < 0 || b >= n {
			return false
		}
		if M[a][b] == 0 {
			return false
		}
		if mark[b] {
			return false
		}
		mark[b] = true
		for i := 0; i < n; i++ {
			dfs(b, i)
		}
		return true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j) {
				circle++
			}
		}
	}
	return circle
}
