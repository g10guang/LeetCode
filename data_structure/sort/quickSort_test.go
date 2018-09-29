package sort

import (
	"testing"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	length := 100
	arr := make([]int, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range arr {
		arr[i] = r.Intn(100)
	}
	QuickSort(arr)
	//sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		assert.True(t, arr[i-1] <= arr[i])
	}
}

func TestQuickSort2(t *testing.T) {
	for i := 0; i < 10000; i++ {
		TestQuickSort(t)
	}
}

func TestQuickSortList(t *testing.T) {
	root := &Node{v: 10, k: 10, next: nil}
	p := root
	p.next = &Node{v: 1, k: 1, next: nil}
	p = p.next
	p.next = &Node{v: 3, k: 3, next: nil}
	p = p.next
	p.next = &Node{v: 13, k: 13, next: nil}
	p = p.next
	printList(t, root)
	t.Log("\n\n")
	QuickSortList(root, nil)
	check(t, root)
}

func TestQuickSortList_OneElem(t *testing.T) {
	root := &Node{v: 10, k: 10, next: nil}
	printList(t, root)
	t.Log("\n\n")
	QuickSortList(root, nil)
	check(t, root)
}

func TestQuickSortList_Random(t *testing.T) {
	for i := 0; i < 100000; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		length := r.Intn(100)
		root := &Node{v: r.Intn(100), k: r.Intn(100), next: nil}
		p := root
		for i := 1; i < length; i++ {
			v := r.Intn(100)
			p.next = &Node{v: v, k: v, next: nil}
			p = p.next
		}
		//printList(t, root)
		//t.Log("\n\n")
		QuickSortList(root, nil)
		check(t, root)
	}
}

func printList(t *testing.T, root *Node) {
	p := root
	for p != nil {
		t.Logf("%v\n", p.k)
		p = p.next
	}
}

func check(t *testing.T, root *Node) {
	if root == nil {
		return
	}
	p, q := root, root.next
	for q != nil {
		assert.True(t, q.k >= p.k)
		p = q
		q = q.next
	}
}
