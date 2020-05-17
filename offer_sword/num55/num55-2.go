package main

import "fmt"

// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, result := balance(root)
	return result
}

func balance(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	leftDepth, leftBalance := balance(node.Left)
	rightDepth, rightBalance := balance(node.Right)
	fmt.Printf("node=%v left=%v %v right=%v %v\n", node.Val, leftDepth, leftBalance, rightDepth, rightBalance)
	return maxVal(leftDepth, rightDepth) + 1, leftBalance && rightBalance && (leftDepth-rightDepth <= 1 && leftDepth-rightDepth >= -1)
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
