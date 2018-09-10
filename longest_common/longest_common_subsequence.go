package longest_common

// Use dynamic program to solve longest common subsequence.
// https://zh.wikipedia.org/wiki/%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E5%BA%8F%E5%88%97
// Subsequence maybe separate with each other.
// Substring must be consecutive.
func longestCommonSubsequence(s, t string) int {
	sr, tr := []rune(s), []rune(t)

	same := func(a, b rune) int {
		if a == b {
			return 1
		}
		return 0
	}

	max := func(x, y, z int) int {
		if x < y {
			x = y
		}
		if x < z {
			x = z
		}
		return x
	}

	maxLen := make([][]int, len(sr)+1)
	for i := 0; i < len(sr)+1; i++ {
		maxLen[i] = make([]int, len(tr)+1)
	}

	for x, xr := range sr {
		x += 1
		for y, yr := range tr {
			y += 1
			maxLen[x][y] = max(maxLen[x-1][y-1]+same(xr, yr), maxLen[x-1][y], maxLen[x][y-1])
		}
	}

	return maxLen[len(sr)][len(tr)]
}
