package main

// https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof/
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	slen := len(s)
	plen := len(p)
	isAny := func(j int) bool {
		return j+1 < len(p) && p[j+1] == '*'
	}
	var match func(i, j int) bool
	match = func(i, j int) bool {
		if j == plen {
			return i == slen
		}
		if isAny(j) {
			if match(i, j+2) {
				return true
			}
		}
		if matchChar(s[i], p[j]) {
			if isAny(j) {
				if match(i+1, j) || match(i+1, j+2) {
					return true
				}
			} else {
				return match(i+1, j+1)
			}
		}
		return false
	}
	return match(0, 0)
}

func matchChar(source, target byte) bool {
	if target == '.' {
		return true
	}
	return source == target
}
