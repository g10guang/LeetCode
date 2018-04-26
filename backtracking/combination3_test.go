package backtracking_test

import (
	"testing"
	"github.com/g10guang/leetcode/backtracking"
)

func TestCombinationSum3(t *testing.T) {
	suits := []struct{
		k, n int
		expect [][]int
	} {
		{k:3,n:7,expect:[][]int{{1,2,4}}},
	}
	for _, s := range suits {
		if r := backtracking.CombinationSum3(s.k, s.n); !compare2Dslice(r, s.expect) {
			t.Errorf("k=%v n=%v\n expect=%v\n return=%v\n", s.k, s.n, s.expect, r)
		}
	}
}
