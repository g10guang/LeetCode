package string_

// KMP algorithm to search the substring s is whether in source string t.
// http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html

func kmp(t, s string) bool {
	sl, tl := len(s), len(t)
	if sl > tl {
		return false
	} else if sl == tl {
		return s == t
	}
	partial := partialMatch(s)
	isMatch := false
	for x, xMove, yMove := 0, 0, 0; x < tl && !isMatch; x += xMove {
		isMatch = true
		for y := yMove; y < sl; y++ {
			if t[x+y] != s[y] {
				if y == 0 {
					xMove, yMove = 1, 0
				} else {
					xMove, yMove = y-partial[y-1], partial[y-1]
				}
				isMatch = false
				break
			}
		}
	}
	return isMatch
}

// Calculate the max common length in substrings of s.
func partialMatch(s string) (ret []int) {
	length := len(s)
	ret = make([]int, length)
	for i := 1; i < length-1; i++ {
		for j := i; j > 0; j-- {
			if s[:j] == s[i+1-j:i+1] {
				ret[i] = j
				break
			}
		}
	}
	return
}
