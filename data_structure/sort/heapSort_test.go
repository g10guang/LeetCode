package sort

import (
	"testing"
	"math/rand"
	"time"
)

func TestHeapSortBuildLoop(t *testing.T) {
	for i := 0; i < 1000; i++ {
		TestHeapSortBuild(t)
	}
}

func TestHeapSortBuild(t *testing.T) {
	var h heap
	for i := 1; i < 100; i++ {
		h = append(h, intElem{data: i})
	}
	h.build()
	isHeap(t, h)
}

func isHeap(t *testing.T, h heap) {
	for i := 0; i < (h.len()-1)/2; i++ {
		left := i*2 + 1
		right := left + 1
		if left < h.len() && h[i].compare(h[left]) < 0 {
			t.Errorf("violate heap law. h[%d]=%d h[%d]=%d\n", i, h[i], left, h[left])
		}
		if right < h.len() && h[i].compare(h[right]) < 0 {
			t.Errorf("viorlate heap law. h[%d]=%d h[%d]=%d\n", i, h[i], right, h[right])
		}
	}
}

func TestHeapPush(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var h heap
	var hs heapSort = &h
	for i := 0; i < 1000; i++ {
		hs.push(intElem{r.Int()})
	}
	isHeap(t, h)
}

func TestHeapSortLoop(t *testing.T) {
	for i := 0; i < 1000; i++ {
		TestHeapSort(t)
	}
}

func TestHeapSort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var h heap
	var hs heapSort = &h
	for i := 0; i < 1000; i++ {
		hs.push(intElem{r.Int()})
	}

	isHeap(t, h)

	h.sort()

	isSort(t, h)
}

func isSort(t *testing.T, h heap) {
	for i := 1; i < h.len(); i++ {
		if h[i].compare(h[i-1]) < 0 {
			t.Errorf("heap is not ascendence order.\n")
		}
	}
}
