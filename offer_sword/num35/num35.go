package main

// https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mark := make(map[*Node]*Node)
	if head == nil {
		return nil
	}
	newHead := copyNode(head)
	mark[head] = newHead
	// basically set Next pointer
	for p, k := head.Next, newHead; p != nil; p = p.Next {
		t := copyNode(p)
		k.Next = t
		k = t
		mark[p] = t
	}
	// set random pointer
	for p, k := head, newHead; p != nil; p, k = p.Next, k.Next {
		ptr := mark[p.Random]
		k.Random = ptr
	}
	return newHead
}

func copyNode(n *Node) *Node {
	if n == nil {
		return nil
	}
	return &Node{
		Val: n.Val,
	}
}
