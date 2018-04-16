package alg

import "errors"

type NoLockQueue struct {
	elems   []interface{}
	head    int
	tail    int
	MaxSize int
	sizePlusOne int
}

func (q *NoLockQueue) Len() int {
	return (q.tail - q.head + q.sizePlusOne) % q.sizePlusOne
}

func (q *NoLockQueue) init() {
	// 这里循环队列采用少用一位方法
	q.sizePlusOne = q.MaxSize + 1
	q.elems = make([]interface{}, q.sizePlusOne)
}

func (q *NoLockQueue) Enqueue(x interface{}) error {
	if q.elems == nil {
		q.init()
	}
	if q.Len() == q.MaxSize {
		return errors.New("full queue")
	}
	q.elems[q.tail] = x
	q.tail = (q.tail + 1) % q.sizePlusOne
	return nil
}

func (q *NoLockQueue) Dequeue() (interface{}, error) {
	if q.Len() == 0 {
		return nil, errors.New("empty queue")
	}
	x := q.elems[q.head]
	q.head = (q.head + 1) % q.sizePlusOne
	return x, nil
}
