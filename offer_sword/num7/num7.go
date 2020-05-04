package main

import "fmt"

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var root *TreeNode
	if len(preorder) != len(inorder) {
		panic("length mismatch")
	}
	if len(preorder) == 0 || len(inorder) == 0 {
		return root
	}
	root = &TreeNode{}
	v := preorder[0]
	index := search(inorder, v)
	if index == -1 {
		panic(fmt.Sprintf("binarySearch not found inorder=%+v v=%v", inorder, v))
	}
	root.Val = v
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}

func search(array []int, target int) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}
	return -1
}
