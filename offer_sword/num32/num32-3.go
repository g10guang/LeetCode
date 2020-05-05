package main

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queueItem struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	var queue []*queueItem
	var ret [][]int
	queue = append(queue, &queueItem{node: root, level: 0})
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if item.node == nil {
			continue
		}
		for len(ret) <= item.level {
			ret = append(ret, make([]int, 0))
		}
		ret[item.level] = append(ret[item.level], item.node.Val)
		queue = append(queue, &queueItem{
			node:  item.node.Left,
			level: item.level + 1,
		}, &queueItem{
			node:  item.node.Right,
			level: item.level + 1,
		})
	}
	for i := 1; i < len(ret); i += 2 {
		revertSlice(ret[i])
	}
	return ret
}

func revertSlice(s []int) {
	for low, high := 0, len(s)-1; low < high; low, high = low+1, high-1 {
		s[low], s[high] = s[high], s[low]
	}
}
