package Huffman

import (
	"github.com/g10guang/leetcode/data_structure/priorityQueue"
	"container/heap"
)

// Node in huffman tree
type Node struct {
	Left   *Node
	Right  *Node
	Weight int
	Value  interface{}
}

type Value struct {
	Weight int
	Value  interface{}
}

type Tree struct {
	Root *Node
	pq   priorityQueue.PriorityQueue
}

// push elem into priority queue
func (t *Tree) Init(data []*Value) {
	heap.Init(&t.pq)
	for _, v := range data {
		heap.Push(&t.pq, &priorityQueue.Item{Value: &Node{Weight: v.Weight, Value: v.Value}, Priority: v.Weight, Index: -1})
	}
}

// Build huffman tree
func (t *Tree) Build() bool {
	for t.pq.Len() > 1 {
		v1 := heap.Pop(&t.pq)
		v2 := heap.Pop(&t.pq)
		item1, ok := v1.(*priorityQueue.Item)
		assertTrue(ok)
		item2, ok := v2.(*priorityQueue.Item)
		assertTrue(ok)
		n1, ok := item1.Value.(*Node)
		assertTrue(ok)
		n2, ok := item2.Value.(*Node)
		assertTrue(ok)
		newNode := &Node{Left: n1, Right: n2, Weight: n1.Weight + n2.Weight, Value: nil}
		heap.Push(&t.pq, &priorityQueue.Item{Value: newNode, Priority: newNode.Weight})
	}
	v := t.pq.Pop()
	item, ok := v.(*priorityQueue.Item)
	assertTrue(ok)
	t.Root, ok = item.Value.(*Node)
	assertTrue(ok)
	return ok
}

func assertTrue(ok bool) {
	if !ok {
		panic("Error Type\n")
	}
}
