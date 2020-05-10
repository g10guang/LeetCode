package main

// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var ret [][]int
	var path []int
	var acc int
	var fn func(node *TreeNode)
	fn = func(node *TreeNode) {
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
			acc -= node.Val
		}()
		acc += node.Val
		if node.Left == nil && node.Right == nil {
			if acc == sum {
				t := make([]int, len(path))
				copy(t, path)
				ret = append(ret, t)
			}
		}
		if node.Left != nil {
			fn(node.Left)
		}
		if node.Right != nil {
			fn(node.Right)
		}
	}
	if root == nil {
		return ret
	}
	fn(root)
	return ret
}
