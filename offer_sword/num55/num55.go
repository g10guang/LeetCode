package main

// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var reserve func(node *TreeNode) int
	reserve = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(reserve(node.Left), reserve(node.Right)) + 1
	}
	return reserve(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}