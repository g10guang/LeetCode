package string_

import (
	"sort"
)

// Note: Suppose all character is ascii.

// Use suffix array.
func longestCommonSubstring(s, t string) string {
	// merge to one string and use suffix array.
	m := s + t
	sl, ml := len(s), len(m)
	arr := make([]int, ml)
	for i := 0; i < ml; i++ {
		arr[i] = i
	}
	maxLen, start := 0, 0
	sort.Sort(&suffixArray{arr, m})
	for x := 0; x < ml-1; x++ {
		ax := arr[x]
		for y := x + 1; y < ml; y++ {
			// count common part.
			ay := arr[y]
			if ax > ay {
				ax, ay = ay, ax
			}
			if ax >= sl && ay < sl {
				continue
			}
			if v := commLen(ax, ay, m); v > maxLen {
				if ax+v > sl {
					v = sl - ax
				}
				if v > maxLen {
					maxLen, start = v, ax
				}
			}
		}
	}
	return s[start : start+maxLen]
}

// Find longest common substring in a string
func longestCommonSubtringInAString(s string) string {
	length := len(s)
	indexArr := make([]int, length)
	for i := 0; i < length; i++ {
		indexArr[i] = i
	}
	sa := &suffixArray{indexArr, s}
	sort.Sort(sa)
	maxLen, start := 0, 0
	for x := 0; x < length-1; x++ {
		for y := x + 1; y < length; y++ {
			if tmp := commLen(x, y, s); tmp > maxLen {
				maxLen = tmp
				start = x
			}
		}
	}
	return s[start : start+maxLen]
}

// Find substrings which appears at least M times.
func substringAtLeastMTimes(s string, M int) (ret []string) {
	length := len(s)
	indexArr := make([]int, length)
	for i := 0; i < length; i++ {
		indexArr[i] = i
	}
	sa := &suffixArray{indexArr, s}
	sort.Sort(sa)
	for x := 0; x+M-1 < length; x++ {
		if v := commLen(indexArr[x], indexArr[x+M-1], s); v > 0 {
			if l := len(ret); l == 0 || s[indexArr[x]:indexArr[x]+v] != ret[l-1] {
				ret = append(ret, s[indexArr[x]:indexArr[x]+v])
			}
		}
	}
	return
}

func commLen(x, y int, s string) (cnt int) {
	if x > y {
		x, y = y, x
	}
	for y < len(s) && s[x] == s[y] {
		cnt++
		x++
		y++
	}
	return cnt
}

type suffixArray struct {
	indexArr []int
	s        string
}

func (o *suffixArray) Len() int {
	return len(o.indexArr)
}

func (o *suffixArray) Less(x, y int) bool {
	return o.s[o.indexArr[x]:] < o.s[o.indexArr[y]:]
}

func (o *suffixArray) Swap(x, y int) {
	o.indexArr[x], o.indexArr[y] = o.indexArr[y], o.indexArr[x]
}
