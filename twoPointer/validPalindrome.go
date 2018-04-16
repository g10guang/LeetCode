package twoPointer

func validPalindrome(s string) bool {
	hasChance := false
	p, q := 0, len(s)-1
	for p < q {
		if s[p] != s[q] {
			if s[p] == s[q-1] {
				// delete q
				hasChance = isPalindrome(s[p:q])
				if hasChance {
					return true
				}
			}
			if s[q] == s[p+1] {
				// deltete q
				hasChance = isPalindrome(s[p+1 : q+1])
				if hasChance {
					return true
				}
			}
			return false
		} else {
			p++
			q--
		}
	}
	return true
}

// no any chance to delete a char
func isPalindrome(s string) bool {
	p, q := 0, len(s)-1
	for p < q {
		if s[p] != s[q] {
			return false
		}
		p++
		q--
	}
	return true
}
