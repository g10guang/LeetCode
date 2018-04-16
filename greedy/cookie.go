package greedy

import "sort"

type S struct {
	t []int
}

// Len() int
// // Less reports whether the element with
// // index i should sort before the element with index j.
// Less(i, j int) bool
// // Swap swaps the elements with indexes i and j.
// Swap(i, j int)

func (s S) Len() int {
	return len(s.t)
}

func (s S) Less(i, j int) bool {
	return s.t[i] < s.t[j]
}

func (s S) Swap(i, j int) {
	s.t[i], s.t[j] = s.t[j], s.t[i]
}

func findContentChildren(g []int, s []int) int {
	sort.Sort(S{t: g})
	sort.Sort(S{t: s})
	satisfy := 0
	gh := len(g) - 1
	for sh := len(s) - 1; sh >= 0; sh-- {
		if s[sh] >= g[gh] {
			satisfy++
		}
		gh--
		if gh < 0 {
			break
		}
		sh++
	}
	return satisfy
}

var FindContentChildren = findContentChildren
