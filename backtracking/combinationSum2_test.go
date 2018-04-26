package backtracking_test

import (
	"testing"
	"github.com/g10guang/leetcode/backtracking"
)

func TestCombinationSum2(t *testing.T) {
	suits := []struct{
		candidates []int
		target int
		expect [][]int
	} {
		{candidates:[]int{10,1,2,7,6,1,5},target:8,expect:[][]int{{1,1,6},{1,2,5},{1,7},{2,6}}},
	}
	for _, s := range suits {
		if r := backtracking.CombinationSum2(s.candidates, s.target); !compare2Dslice(s.expect, r) {
			t.Errorf("candidates=%v\n target=%v\n expect=%v\n return=%v\n", s.candidates, s.target, s.expect, r)
		}
	}
}