package twoPointer

// https://leetcode.com/problems/reverse-vowels-of-a-string/description/
func reverseVowels(s string) string {
	target := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	slice := make([]byte, len(s))
	for p, q := 0, len(s)-1; p <= q; {
		for p < len(s) && !target[s[p]] {
			slice[p] = s[p]
			p++
		}
		for q >= 0 && !target[s[q]] {
			slice[q] = s[q]
			q--
		}
		if p < q {
			slice[p] = s[q]
			slice[q] = s[p]
		} else if p == q {
			slice[p] = s[p]
		}
		p++
		q--
	}
	return string(slice)
}

var ReverseVowels = reverseVowels