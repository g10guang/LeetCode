package greedy_test

import (
	"testing"
	"github.com/g10guang/leetcode/greedy"
)

func TestReconstructQueue(t *testing.T) {
	suits := []struct{
		input [][]int
		expect [][]int
	} {
		{[][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}, [][]int{{5,0}, {7,0}, {5,2}, {6,1}, {4,4}, {7,1}}},
	}
	for _, s := range suits {
		if r := greedy.ReconstructQueue(s.input); !compare2Dslice(r, s.expect) {
			t.Errorf("\n input=%v\n expect=%v\n return=%v\n", s.input, s.expect, r)
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