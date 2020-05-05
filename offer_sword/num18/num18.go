package main

// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	for parent, current := head, head.Next; current != nil; parent, current = current, current.Next {
		if current.Val == val {
			parent.Next = current.Next
			break
		}
	}
	return head
}
