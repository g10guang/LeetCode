package sort

// big root heap
type heapSort interface {
	// swim the heap[pos] and return the right position of heap[pos]
	swim(int) int

	// sink the heap[pos] and return the right position of heap[pos]
	sink(int) int

	// push a new element to heap
	// return the position of x in array.
	push(comparable) int

	// return the top element in heap
	// if len() == 0: panic
	top() comparable

	// return the length of heap
	len() int

	// update the position of element
	// return the right position.
	update(int) int

	// with heap have more than one elem.
	// make data order to be heap.
	build()

	// make heap elements to be the ascendant order.
	sort()
}

// store the heap into an 1-index array.
// for element in x index.
// parent(x) = x / 2
// left(x) = x * 2
// right(x) = x * 2 + 1

type comparable interface {
	// cur > o return 1
	// cur == o return 0
	// cur < o return -1
	// type(cur) != type(o) panic("type error")
	compare(o comparable) int
}

type intElem struct {
	data int
}

func (o intElem) compare(t comparable) int {
	m, ok := t.(intElem)
	if !ok {
		panic("Type Error")
	}
	switch {
	case o.data > m.data:
		return 1
	case o.data < m.data:
		return -1
	default:
		return 0
	}
}

type heap []comparable

func (o *heap) len() int {
	return len(*o)
}

func (o *heap) top() comparable {
	if o.len() > 0 {
		return (*o)[0]
	}
	panic("Len is zero.")
}

func (o *heap) push(x comparable) int {
	*o = append(*o, x)
	return o.update(len(*o) - 1)
}

func (o *heap) pop() comparable {
	if o.len() < 0 {
		panic("Len is Zero.")
	}
	// pop the roof of current heap.
	ret := (*o)[0]
	(*o)[0] = (*o)[o.len()-1]
	*o = (*o)[:o.len()-1]
	if o.len() > 0 {
		// after pop, maybe o.len() == 0
		o.sink(0)
	}
	return ret
}

func (o *heap) update(pos int) int {
	checkIndex(pos, o.len())

	if t := o.swim(pos); t != pos {
		return t
	}

	return o.sink(pos)
}

func (o *heap) swim(pos int) int {
	checkIndex(pos, o.len())
	parent := (pos - 1) / 2
	if parent >= 0 && (*o)[pos].compare((*o)[parent]) > 0 {
		(*o)[pos], (*o)[parent] = (*o)[parent], (*o)[pos]
		return o.swim(parent)
	}
	// pos is the right position.
	return pos
}

func (o *heap) sink(pos int) int {
	checkIndex(pos, o.len())
	left := pos*2 + 1
	right := left + 1
	max := pos
	if left < o.len() && (*o)[max].compare((*o)[left]) < 0 {
		max = left
	}
	if right < o.len() && (*o)[max].compare((*o)[right]) < 0 {
		max = right
	}
	if max != pos {
		(*o)[max], (*o)[pos] = (*o)[pos], (*o)[max]
		return o.sink(max)
	}
	// pos is the right position.
	return pos
}

func (o *heap) build() {
	// swim(o[pos]) for pos in range(1, len(o))
	for i := 0; i < o.len(); i++ {
		o.swim(i)
	}
}

func checkIndex(pos, length int) {
	if pos < 0 || pos >= length {
		panic("Index out of range.")
	}
}

func (o *heap) sort() {
	t := *o
	for i := t.len() - 1; i >= 0; i-- {
		(*o)[i] = t.pop()
	}
}
