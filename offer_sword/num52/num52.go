package main

import "fmt"

// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	la, lb := length(headA), length(headB)
	large, small := headB, headA
	if la > lb {
		large, small = headA, headB
	}
	more := la - lb
	if more < 0 {
		more = -more
	}
	fmt.Printf("la=%v lb=%v large=%v small=%v\n", la, lb, large.Val, small.Val)
	for i := 0; i < more; i++ {
		large = large.Next
	}
	fmt.Printf("large=%v small=%v\n", large.Val, small.Val)
	for large != nil && small != nil {
		if large == small {
			return large
		}
		large = large.Next
		small = small.Next
	}
	return nil
}

func length(root *ListNode) int {
	p := root
	result := 0
	for p != nil {
		result++
		p = p.Next
	}
	return result
}

func main() {
	common := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	headA := &ListNode{
		Val:  1,
		Next: common,
	}

	headB := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: common,
		},
	}

	fmt.Printf("%+v\n", getIntersectionNode(headA, headB))
}
