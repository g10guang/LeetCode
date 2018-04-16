package graph_test

import (
	"testing"
	"github.com/g10guang/leetcode/graph"
)

func TestSolve(t *testing.T) {
	suits := []struct{
		input [][]byte
		expect [][]byte
	} {
		{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}, [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		{[][]byte{{'X','O','X','X'},{'O','X','O','X'},{'X','O','X','O'},{'O','X','O','X'},{'X','O','X','O'},{'O','X','O','X'}}, [][]byte{{'X','O','X','X'},{'O','X','X','X'},{'X','X','X','O'},{'O','X','X','X'},{'X','X','X','O'},{'O','X','O','X'}}},
		//{[][]byte{}, [][]byte{}},
		//{[][]byte{{'O','O','O','O','X','X'},{'O','O','O','O','O','O'},{'O','X','O','X','O','O'},{'O','X','O','O','X','O'},{'O','X','O','X','O','O'},{'O','X','O','O','O','O'}}, [][]byte{}},
		{[][]byte{{'X','X','X','X','X'},{'X','O','O','O','X'},{'X','X','O','O','X'},{'X','X','X','O','X'},{'X','O','X','X','X'}}, [][]byte{{'X','X','X','X','X'},{'X','X','X','X','X'},{'X','X','X','X','X'},{'X','X','X','X','X'},{'X','O','X','X','X'}}},
	}
	for _, s := range suits {
		t.Logf("input=%v\n", s.input)
		graph.Solve(s.input)
		if !compareByteSlice(s.expect, s.input) {
			t.Errorf("expect=%v return=%v", s.expect, s.input)
		}
	}
}

func compareByteSlice(s, t [][]byte) bool {
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if len(v) != len(t[i]) {
			return false
		}
		for j, l := range v {
			if l != t[i][j] {
				return false
			}
		}
	}
	return true
}