package main

// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	first := l1
	second := l2
	var head *ListNode
	if first == nil {
		return second
	} else if second == nil {
		return first
	} else if first.Val <= second.Val {
		head = first
		first = first.Next
	} else {
		head = second
		second = second.Next
	}

	current := head
	for first != nil && second != nil {
		if first.Val <= second.Val {
			current.Next = first
			current = first
			first = first.Next
		} else {
			current.Next = second
			current = second
			second = second.Next
		}
	}

	for first != nil {
		current.Next = first
		current = first
		first = first.Next
	}

	for second != nil {
		current.Next = second
		current = second
		second = second.Next
	}
	return head
}
