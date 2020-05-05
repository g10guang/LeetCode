package main

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	parent := head
	current := parent.Next
	head.Next = nil
	for current != nil {
		t := current.Next
		current.Next = parent
		parent = current
		current = t
	}
	return parent
}
