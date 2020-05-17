package main

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var queue []*TreeNode

func levelOrder(root *TreeNode) []int {
	var ret []int
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		queue = append(queue, node.Left, node.Right)
		ret = append(ret, node.Val)
	}
	return ret
}
