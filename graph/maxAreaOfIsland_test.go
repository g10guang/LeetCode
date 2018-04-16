package graph_test

import (
	"testing"
	"github.com/g10guang/leetcode/graph"
)

func TestMaxAreaOfIsland(t *testing.T) {
	suits := []struct{
		input [][]int
		expect int
	} {
		{[][]int{{1,1,0,0,0}, {1,1,0,0,0}, {0,0,0,1,1}, {0,0,0,1,1}}, 4},
		{[][]int{{0,1,1}}, 2},
	}
	for _, s := range suits {
		if r := graph.MaxAreaOfIsland(s.input); r != s.expect {
			t.Errorf("input=%v expect=%v return=%v", s.input, s.expect, r)
		}
	}
}
