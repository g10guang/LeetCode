package graph_test

import (
	"testing"
	"github.com/g10guang/leetcode/graph"
)

func TestBinaryTreePaths(t *testing.T) {
	root := &graph.TreeNode{Val:1}
	B := &graph.TreeNode{Val:2}
	C := &graph.TreeNode{Val:3}
	D := &graph.TreeNode{Val:4}
	E := &graph.TreeNode{Val:5}
	F := &graph.TreeNode{Val:6}
	G := &graph.TreeNode{Val:7}
	root.Left = B
	root.Right = C
	B.Left = D
	B.Right = E
	C.Left = F
	C.Right = G
	suits := []struct{
		input *graph.TreeNode
		expect []string
	} {
		{root, []string{"1->2->4", "1->2->5", "1->3->6", "1->3->7"}},
		{nil, []string{}},
	}
	for _, s := range suits {
		if r := graph.BinaryTreePaths(s.input); !compare(s.expect, r) {
			t.Errorf("input=%v expect=%v return=%v", s.input, s.expect, r)
		}
	}
}

func compare(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if v != t[i] {
			return false
		}
	}
	return true
}
