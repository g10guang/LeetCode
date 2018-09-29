package sort

import (
	"math/rand"
	"time"
)

func QuickSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	// scratch a random number.
	// expect to choose a number to split the array averagely.
	//pivot := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(length)
	//pivot := 0
	//arr[pivot], arr[0] = arr[0], arr[pivot]
	split := arr[0]
	splitIndex := 0
	scanner := 1
	last := length - 1
	for scanner <= last {
		if arr[scanner] < split {
			arr[splitIndex], arr[scanner] = arr[scanner], arr[splitIndex]
			splitIndex = scanner
			scanner++
		} else if arr[scanner] > split {
			arr[last], arr[scanner] = arr[scanner], arr[last]
			last--
		} else {
			scanner++
		}
	}
	// arr[:splitIndex] <= split
	// arr[splitIndex+1:] > split
	// unfortunately split is the max value of arr. every recurse will just reduce one element.
	QuickSort(arr[:splitIndex])
	QuickSort(arr[splitIndex+1:])
}

func QuickSort2(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	// scratch a random number.
	// expect to choose a number to split the array averagely.
	pivot := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(length)
	arr[pivot], arr[0] = arr[0], arr[pivot]
	split := arr[0]
	scanner := 1
	last := length - 1
	for scanner <= last {
		if arr[scanner] <= split {
			scanner++
		} else {
			arr[last], arr[scanner] = arr[scanner], arr[last]
			last--
		}
	}
	arr[scanner-1], arr[0] = arr[0], arr[scanner-1]
	QuickSort(arr[:scanner-1])
	QuickSort(arr[scanner:])
}

// choose the last element as the pivot.
// low point to the boundary of elem <= pivot.
func QuickSort3(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	pivot := arr[length-1]
	low := -1
	for i := 0; i < length-1; i++ {
		// move the elem bigger than pivot to the right side.
		if arr[i] <= pivot {
			low++
			arr[i], arr[low] = arr[low], arr[i]
		}
	}
	// swap the pivot to the mid position.
	// low + 1 point to elem > pivot.
	arr[length-1], arr[low+1] = arr[low+1], arr[length-1]
	QuickSort3(arr[:low])
	QuickSort3(arr[low+1:])
}

type Node struct {
	v    interface{}
	k    int
	next *Node
}

// Quick sort for list [start, end)
func QuickSortList(start, end *Node) {
	if start == end {
		return
	}
	p, q := start, start.next
	for q != end {
		if q.k < p.k {
			// swap the value, not change list structure
			swap(p, q)
			p = p.next
			swap(p, q)
		}
		q = q.next
	}
	// sort [start, p)
	QuickSortList(start, p)
	// sort (p, end)
	QuickSortList(p.next, end)
}

func swap(p, q *Node) {
	p.v, q.v = q.v, p.v
	p.k, q.k = q.k, p.k
}
