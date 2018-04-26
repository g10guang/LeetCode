package backtracking_test

import (
	"testing"
	"github.com/g10guang/leetcode/backtracking"
)

func TestCombinationSum(t *testing.T) {
	suits := []struct{
		candidates []int
		target int
		expect [][]int
	} {
		{candidates:[]int{2,3,6,7}, target:7, expect:[][]int{{2,2,3},{7}}},
	}
	for _, s := range suits {
		if r := backtracking.CombinationSum(s.candidates, s.target); !compare2Dslice(s.expect, r) {
			t.Errorf("candidates=%v\n target=%v\n expect=%v\n return=%v\n", s.candidates, s.target, s.expect, r)
		}
	}
}
