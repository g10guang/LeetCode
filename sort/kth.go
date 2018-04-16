package sort

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[j], h[i] = h[i], h[j]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

// Use heap sort to find the kth largest element.
// https://leetcode.com/problems/kth-largest-element-in-an-array/description/
func HeapFindKthLargest(data []int, k int) int {
	// assume k is valid.
	h := make(IntHeap, k)
	f := true
	for i, v := range data{
		if i < k {
			h[i] = v
		} else if f {
			heap.Init(&h)
			f = false
		}
		if !f && h[0] < v{
			h[0] = v
			heap.Fix(&h, 0)
		}
	}
	if f {
		heap.Init(&h)
	}
	return h[0]
}

// Use Quick Sort to realize
func QsortFindKthLargest(data []int, k int) int {
	t := len(data) - k
	l, u := 0, len(data) - 1
	for mid := qsort(data, l, u); mid != t; mid = qsort(data, l, u){
		if mid < t {
			l = mid + 1
		} else {
			u = mid - 1
		}
	}
	return data[t]
}

// One time quick sort [l,u]
func qsort(data []int, l, u int) int {
	midElem := data[l]
	head := l
	for i, tail := l + 1, u; i <= tail; {
		if data[i] > midElem {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			i++
			head++
		}
	}
	data[head] = midElem
	return head
}