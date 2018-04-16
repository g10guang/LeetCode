package alg

import "errors"

type listNode struct {
	value interface{}
	next *listNode
}

type listStack struct {
	root *listNode
	length int
}

func (l *listStack) Len() int {
	return l.length
}

func (l *listStack) Top() (interface{}, error) {
	if l.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	return l.root.value, nil
}

func (l *listStack) IsEmpty() bool {
	return l.Len() == 0
}

func (l *listStack) Push(x interface{}) {
	newNode := &listNode{value:x, next:l.root}
	l.root = newNode
	l.length++
}

func (l *listStack) Pop() (interface{}, error) {
	if l.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	t := l.root
	l.root = l.root.next
	t.next = nil
	l.length--
	return t, nil
}
