package main

// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return root.Left == root.Right
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(A, B *TreeNode) bool {
	if A == nil || B == nil {
		return A == B
	}
	if A.Val != B.Val {
		return false
	}
	return isMirror(A.Left, B.Right) && isMirror(A.Right, B.Left)
}
