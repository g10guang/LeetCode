package main

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var ret *ListNode
	var fn func(node *ListNode)
	fn = func(node *ListNode) {
		if node == nil {
			return
		}
		fn(node.Next)
		if k--; k == 0 {
			ret = node
		}
	}
	fn(head)
	return ret
}