package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	var ret []int
	var fn func(node *ListNode)
	fn = func(node *ListNode) {
		if node == nil {
			return
		}
		fn(node.Next)
		ret = append(ret, node.Val)
	}
	fn(head)
	return ret
}
