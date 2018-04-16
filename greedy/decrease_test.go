package greedy_test

import (
	"testing"

	"github.com/g10guang/leetcode/greedy"
)

func TestDesease(t *testing.T) {
	suits := []struct {
		input    []int
		expected bool
	}{
		{[]int{4, 2, 3}, true},
		{[]int{1, 3, 2}, true},
		{[]int{3, 4, 2, 3}, false},
	}
	for _, suit := range suits {
		if r := greedy.CheckPossibility2(suit.input); r != suit.expected {
			t.Errorf("input=%v expected=%v return=%v", suit.input, suit.expected, r)
		}
	}
}
