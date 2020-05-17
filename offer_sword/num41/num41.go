package main

import (
	"container/heap"
	"fmt"
)

type IntHeap struct {
	flag bool // false-->MaxHeap  true-->MinHeap
	data []int
}

func newMaxHeap() *IntHeap {
	return &IntHeap{
		flag: false,
	}
}

func newMinHeap() *IntHeap {
	return &IntHeap{
		flag: true,
	}
}

func (h *IntHeap) Len() int {
	return len(h.data)
}

func (h *IntHeap) Less(i, j int) bool {
	if h.flag {
		return h.data[i] < h.data[j]
	}
	return h.data[i] > h.data[j]
}
func (h *IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *IntHeap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	item := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return item
}

func (h *IntHeap) Top() int {
	item := heap.Pop(h)
	heap.Push(h, item)
	return item.(int)
}

// https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/
type MedianFinder struct {
	maxHeap *IntHeap
	minHeap *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: newMaxHeap(),
		minHeap: newMinHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// add minHeap first
	if this.maxHeap.Len() == this.minHeap.Len() {
		heap.Push(this.maxHeap, num)
		item := heap.Pop(this.maxHeap)
		heap.Push(this.minHeap, item)
	} else {
		heap.Push(this.minHeap, num)
		item := heap.Pop(this.minHeap)
		heap.Push(this.maxHeap, item)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		if this.maxHeap.Len() == 0 {
			return 0
		}
		a := float64(this.maxHeap.Top())
		b := float64(this.minHeap.Top())
		return (a + b) / 2
	}
	return float64(this.minHeap.Top())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func heapTest() {
	maxHeap := newMaxHeap()
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	fmt.Printf("maxHeap.Top=%d\n", maxHeap.Top())
	for maxHeap.Len() > 0 {
		fmt.Printf("maxHeap.Top=%v\n", heap.Pop(maxHeap))
	}

	minHeap := newMinHeap()
	heap.Push(minHeap, 1)
	heap.Push(minHeap, 3)
	heap.Push(minHeap, 2)
	fmt.Printf("minHeap.Top=%d\n", minHeap.Top())
	for minHeap.Len() > 0 {
		fmt.Printf("minHeap.Top=%v\n", heap.Pop(minHeap))
	}
}


func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(3)
	obj.AddNum(2)
	obj.AddNum(4)
	fmt.Printf("1 %v\n", obj.FindMedian())
}