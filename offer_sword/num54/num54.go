package main

import "fmt"

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var reserve func(node *TreeNode) *TreeNode
	reserve = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		rightResult := reserve(node.Right)
		if rightResult != nil {
			return rightResult
		}
		if k--; k == 0 {
			return node
		}
		return reserve(node.Left)
	}
	return reserve(root).Val
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	fmt.Printf("%v\n", kthLargest(root, 1))
}