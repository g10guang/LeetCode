package twoPointer

func findLongestWord(s string, d []string) string {
	maxIndex := -1
	for i, t := range d {
		if isIn(s, t) {
			if maxIndex == -1 {
				maxIndex = i
			} else if isBigger(t, d[maxIndex]) {
				maxIndex = i
			}
		}
	}
	if maxIndex == -1 {
		return ""
	} else {
		return d[maxIndex]
	}
}

func isBigger(t1, t2 string) bool {
	if len(t1) > len(t2) {
		return true
	} else if len(t1) < len(t2) {
		return false
	} else {
		return t1 < t2
	}
}

func isIn(s, t string) bool {
	if len(s) < len(t) {
		return false
	}
	ht := 0
	for hs := 0; hs < len(s) && ht < len(t); hs++ {
		if s[hs] == t[ht] {
			ht++
		}
	}
	return ht >= len(t)
}