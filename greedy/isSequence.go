package greedy

//https://leetcode.com/problems/is-subsequence/description/
func isSubsequence(s string, t string) bool {
	tlen := len(t)
	slen := len(s)
	j := 0
	for i := 0; i < tlen && j < slen; i++ {
		if t[i] == s[j] {
			j++
		}
	}
	return j >= slen
}
