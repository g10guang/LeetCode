package graph_test

import (
	"testing"
	"github.com/g10guang/leetcode/graph"
)

func TestNumIslands(t *testing.T) {
	suits := []struct{
		input [][]byte
		expect int
	} {
		{[][]byte{}, 0},
		{[][]byte{{1,1,1,1,0},{1,1,0,1,0},{1,1,0,0,0},{0,0,0,0,0}}, 1},
		{[][]byte{{1,1,0,0,0},{1,1,0,0,0},{0,0,1,0,0},{0,0,0,1,1}}, 3},
	}
	for _, s := range suits {
		if r := graph.NumIslands(s.input); r != s.expect {
			t.Errorf("input=%v expect=%v return=%v", s.input, s.expect, r)
		} else {
			t.Logf("Success return=%v", r)
		}
	}
}
