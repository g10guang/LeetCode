package binarysearch_test

import (
	"testing"
)

func TestCoin(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 3},
		{9, 3},
		{10, 4},
	}
	for _, suit := range cases {
		if r := arrangeCoins(suit.input); r != suit.expected {
			t.Errorf("n=%d return=%d expected=%d\n", suit.input, r, suit.expected)
		} else {
			t.Logf("n=%d return=%d", suit.input, r)
		}
	}
}
