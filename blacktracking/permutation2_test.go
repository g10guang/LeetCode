package blacktracking_test

import (
	"testing"
	"github.com/g10guang/leetcode/blacktracking"
)

func TestPermute2(t *testing.T) {
	suits := []struct{
		input []int
		expect [][]int
	} {
		//{[]int{1,2,3}, [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}},
		{[]int{1,1,2}, [][]int{{1,1,2},{1,2,1},{2,1,1}}},
	}
	for _, s := range suits {
		if r := blacktracking.PermuteUnique(s.input); !compare2Dslice(r, s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, r)
		}
	}
}
