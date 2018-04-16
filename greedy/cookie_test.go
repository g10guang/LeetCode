package greedy_test

import (
	"testing"
	"github.com/g10guang/leetcode/greedy"
)


func TestCookie(t *testing.T) {
	cases := []struct {
		g        []int
		s        []int
		expected int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
	}
	for _, suit := range cases {
		if r := greedy.FindContentChildren(suit.g, suit.s); r != suit.expected {
			t.Errorf("g=%v s=%v return=%d expected=%d\n", suit.g, suit.s, r, suit.expected)
		}
	}
}
