package sort

import "container/heap"

type Item struct {
	priority int
	value int
}

type PriorityHeap []*Item

func (pq *PriorityHeap) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityHeap) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityHeap) Len() int {
	return len(*pq)
}

func (pq *PriorityHeap) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityHeap) Pop() interface{} {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return x
}

// return the k most frequent elements.
func topKFrequent(nums []int, k int) []int {
	pq := PriorityHeap(make([]*Item, k))
	dict := make(map[int]int)
	for _, v := range nums {
		dict[v] = dict[v] + 1
	}
	i := 0
	for t, v := range dict {
		if i < k {
			pq[i] = &Item{v, t}
		}
		if i == k {
			heap.Init(&pq)
		}
		if pq[0].priority < v {
			pq[0].priority = v
			pq[0].value = t
			heap.Fix(&pq, 0)
		}
		i++
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(&pq).(*Item).value
	}
	for i, j := 0, k-1; i < j; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

var SortTopKFrequent = topKFrequent
