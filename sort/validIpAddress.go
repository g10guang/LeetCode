package sort

import (
	"strconv"
)

// https://leetcode.com/problems/restore-ip-addresses/description/
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	var resolve func(level, start int, path string)
	resolve = func(level, start int, path string) {
		if start >= len(s) {
			return
		}
		if level == 4 {
			if s[start] == '0' && start < len(s)-1 {
				return
			}
			if _, b := between0And255(s[start:]); b {
				result = append(result, path+"."+s[start:])
			}
			return
		}
		t := 1
		for _, isValid := between0And255(s[start : start+t]); t+start < len(s) && isValid; _, isValid = between0And255(s[start : start+t]) {
			if level == 1 {
				resolve(level+1, start+t, s[start:start+t])
			} else {
				resolve(level+1, start+t, path+"."+s[start:start+t])
			}
			if s[start] == '0' {
				break
			}
			t++
		}
		return
	}
	resolve(1, 0, "")
	return result
}

func between0And255(snippet string) (int, bool) {
	t, err := strconv.Atoi(snippet)
	if err != nil {
		return 0, false
	}
	return t, t >= 0 && t <= 255
}

var RestoreIpAddresses = restoreIpAddresses
