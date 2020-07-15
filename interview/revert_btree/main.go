package main

import (
	"encoding/json"
	"fmt"
)

// 题目描述，二叉树翻转
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 具体翻转
func (node *TreeNode) Reverse() {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	if node.Left != nil {
		node.Left.Reverse()
	}
	if node.Right != nil {
		node.Right.Reverse()
	}
}

func (node *TreeNode) Copy() *TreeNode {
	if node == nil {
		return nil
	}
	newNode := &TreeNode{
		Val: node.Val,
	}
	if node.Left != nil {
		newNode.Left = node.Left.Copy()
	}
	if node.Right != nil {
		newNode.Right = node.Right.Copy()
	}
	return newNode
}

func (node *TreeNode) IsMirror(other *TreeNode) bool {
	if node == nil {
		return other == nil
	}
	if other == nil {
		return node == nil
	}
	if node.Val != other.Val {
		return false
	}
	return node.Left.IsMirror(other.Right) && node.Right.IsMirror(other.Left)
}

func (node *TreeNode) String() string {
	b, _ := json.Marshal(node)
	return string(b)
}

func main() {
	test3()
}

func test3() {
	tree := &TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   4,
			},
			Right: nil,
			Val:   2,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Val: 1,
	}
	copyTree := tree.Copy()
	fmt.Printf("OriginTree=%v\n", copyTree.String())
	tree.Reverse()
	fmt.Printf("CopyTree=%v\nreserveTree=%v\n", copyTree.String(), tree.String())
	fmt.Printf("isMirror=%v\n", tree.IsMirror(copyTree))
}

func test2() {
	tree := &TreeNode{
		Left: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   2,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Val: 1,
	}
	copyTree := tree.Copy()
	fmt.Printf("%s\n", copyTree.String())
}

func test1() {
	tree := &TreeNode{
		Left: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   2,
		},
		Right: &TreeNode{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Val: 1,
	}
	tree.Reverse()
	fmt.Printf("%s\n", tree.String())
}
