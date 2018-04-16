package greedy_test

import (
	"testing"

	"github.com/g10guang/leetcode/greedy"
)

func TestStock(t *testing.T) {
	suits := []struct {
		input    []int
		expected int
	}{
		{[]int{8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6}, 15},
		{[]int{1, 2, 3, 4, 2, 6}, 7},
		{[]int{5, 2, 3, 2, 6, 6, 2, 9, 1, 0, 7, 4, 5, 0}, 20},
	}
	for _, suit := range suits {
		if r := greedy.MaxProfit2(suit.input); r != suit.expected {
			t.Errorf("input=%v expected=%v result=%v", suit.input, suit.expected, r)
		}
	}
}
