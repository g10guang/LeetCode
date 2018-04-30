package alg_test

import (
	"testing"
	"github.com/g10guang/leetcode/alg"
)

func TestFindDuplicateNumber(t *testing.T) {
	suits := []struct {
		input  []int
		expect int
	}{
		{input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 7}, expect: 7},
	}
	for _, s := range suits {
		if r := alg.FindDuplicate(s.input); r != s.expect {
			t.Errorf("\n input=%v\n expect=%v\n return=%v", s.input, s.expect, r)
		}
	}
}
