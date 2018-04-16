package binarysearch_test

import (
	"testing"
	"github.com/g10guang/leetcode/binarysearch"
)

func TestMySqrt(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 3},
	}
	for _, suit := range cases {
		if r := binarysearch.MySqrt(suit.input); r != suit.output {
			t.Errorf("input=%d return=%d expected=%d\n", suit.input, r, suit.output)
		} else {
			t.Logf("input=%d return=%d\n", suit.input, r)
		}
	}
}
