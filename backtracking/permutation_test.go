package backtracking_test

import (
	"testing"
	"github.com/g10guang/leetcode/backtracking"
)

func TestPermute(t *testing.T) {
	suits := []struct{
		input []int
		expect [][]int
	} {
		{[]int{1,2,3}, [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}},
	}
	for _, s := range suits {
		if r := backtracking.Permute(s.input); !compare2Dslice(r, s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, r)
		}
	}
}

func compare2Dslice(t, s [][]int) bool {
	if len(t) != len(s) {
		return false
	}
	for i := 0; i < len(t); i++ {
		if len(s[i]) != len(t[i]) {
			return false
		}
		for j := 0; j < len(s[i]); j++ {
			if t[i][j] != s[i][j] {
				return false
			}
		}
	}
	return true
}