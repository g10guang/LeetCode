package graph

import (
	"strings"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/binary-tree-paths/description/
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	curPath := make([]string, 0)
	result := make([]string, 0)
	var handle func(p *TreeNode)
	handle = func(p *TreeNode) {
		if p.Left == nil && p.Right == nil {
			curPath = append(curPath, strconv.Itoa(p.Val))
			result = append(result, strings.Join(curPath, "->"))
		} else {
			curPath = append(curPath, strconv.Itoa(p.Val))
			if p.Left != nil {
				handle(p.Left)
			}
			if p.Right != nil {
				handle(p.Right)
			}
		}
		// Pop out of stack
		curPath = curPath[:len(curPath)-1]
	}
	handle(root)
	return result
}


var BinaryTreePaths = binaryTreePaths
