package graph_test

import (
	"testing"
	"github.com/g10guang/leetcode/graph"
)

func TestPacificAtlantic(t *testing.T) {
	suits := []struct {
		input  [][]int
		expect [][]int
	}{
		{input: [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}, expect: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
	}
	for _, s := range suits {
		if r := graph.PacificAtlantic(s.input); !compareIntMatrix(s.expect, r) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, r)
		}
	}
}

func compareIntMatrix(s, t [][]int) bool {
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if t[i][0] != v[0] || t[i][1] != v[1] {
			return false
		}
	}
	return true
}