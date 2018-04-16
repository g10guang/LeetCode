package greedy_test

import (
	"testing"

	"github.com/g10guang/leetcode/greedy"
)

func TestBalloon(t *testing.T) {
	cases := []struct {
		input  [][]int
		expect int
	}{
		{[][]int{{1, 100}, {4, 500}, {100, 200}, {5, 19}}, 2},
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2147483647}}, 1},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
	}
	for _, suit := range cases {
		if r := greedy.ShotBalloon(suit.input); r != suit.expect {
			t.Errorf("input=%v expected=%v return=%v\n", suit.input, suit.expect, r)
		}
	}
}
