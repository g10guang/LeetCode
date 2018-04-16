package alg

import (
	"errors"
)

// 大根堆
type Heap struct {
	elems   []int
	size    int
	MaxSize int
}

func (h *Heap) isEmpty() bool {
	return h.size == 0
}

func (h *Heap) Size() int {
	return h.size
}

func parent(pos int) int {
	// Go 向 0 取整
	if pos == 0 {
		return -1
	}
	return (pos - 1) / 2
}

func left(pos int) int {
	return 2*pos + 1
}

func right(pos int) int {
	return 2*pos + 2
}

// 元素上浮，重新调整堆
func (h *Heap) Swim(pos int) {
	// pos == 0 已经成为了最大的元素
	for pos > 0 && h.elems[parent(pos)] < h.elems[pos] {
		h.elems[parent(pos)], h.elems[pos] = h.elems[pos], h.elems[parent(pos)]
		pos = parent(pos)
	}
}

// 元素下沉，重新调整堆
func (h *Heap) Sink(pos int) {
	var maxChildPos int
	// 判断该结点有无孩子
	for pos <= parent(h.size-1) {
		//	选出 pos 两个孩子谁大
		if right(pos) >= h.size {
			maxChildPos = left(pos)
		} else {
			if h.elems[left(pos)] < h.elems[right(pos)] {
				maxChildPos = right(pos)
			} else {
				maxChildPos = left(pos)
			}
		}
		if h.elems[pos] < h.elems[maxChildPos] {
			h.elems[pos], h.elems[maxChildPos] = h.elems[maxChildPos], h.elems[pos]
			pos = maxChildPos
		} else {
			break
		}
	}
}

func (h *Heap) init() {
	if h.elems == nil {
		h.elems = make([]int, h.MaxSize)
		h.size = 0
	}
}

// 插入元素到末端，然后对堆进行重新调整
func (h *Heap) Insert(x int) error {
	if h.elems == nil {
		h.init()
	}
	if h.size >= len(h.elems) {
		errors.New("heap is full")
	}
	h.elems[h.size] = x
	h.size++
	h.Swim(h.size - 1)
	return nil
}

func (h *Heap) DelMax() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("head is empty")
	}
	x := h.elems[0]
	h.elems[0] = h.elems[h.size-1]
	h.size--
	h.Sink(0)
	return x, nil
}

// 自底向上构建堆
func (h *Heap) BottomUpConstruct() {
	// 选出第一个有孩子的结点
	for p := parent(h.size - 1); p >= 0; p-- {
		h.Sink(p)
	}
}

// 自顶向下构建堆
func (h *Heap) UpBottomConstruct() {
	if h.size <= 1 {
		return
	}
	for p := 0; p >= parent(h.size-1); p++ {
		h.Swim(p)
	}
}

// 使用堆排序，因为本堆是大根堆，所以得到的排序结构应该是从小到大
func (h *Heap) Sort() {
	for tail := h.size - 1; tail > 0; tail-- {
		h.elems[0], h.elems[tail] = h.elems[tail], h.elems[0]
		h.size --
		h.Sink(0)
	}
}

func (h *Heap) Elems() []int {
	return h.elems
}