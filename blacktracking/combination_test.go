package blacktracking_test

import (
	"testing"
	"github.com/g10guang/leetcode/blacktracking"
)

func TestCombine(t *testing.T) {
	suits := []struct {
		n, k   int
		expect [][]int
	}{
		{4, 2, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
		{5, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}, {3, 4}, {3, 5}, {4, 5},}},
		{1, 1, [][]int{{1}}},
	}
	for _, s := range suits {
		if r := blacktracking.Combine(s.n, s.k); !compare2Dslice(r, s.expect) {
			t.Errorf("n=%v, k=%v\n expect=%v\n return=%v\n", s.n, s.k, s.expect, r)
		}
	}
}
