package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	var fn func(first, second *TreeNode) bool
	fn = func(first, second *TreeNode) bool {
		if second == nil {
			return true
		}
		if first == nil {
			return false
		}
		if first.Val == second.Val {
			if fn(first.Left, second.Left) && fn(first.Right, second.Right) {
				return true
			}
		}
		return fn(first.Left, B) || fn(first.Right, B)
	}
	return fn(A, B)
}
