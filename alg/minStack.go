package alg

import "errors"

/*
增加 O(1) 空间复杂，支持可以取出当前栈最小值
进栈操作：
设要插入的元素为 x
1. 如果队列为空，直接插入，然后更新 min 值
2. 如果队列不为空，则将 x 与 min 做比较，如果 x >= min 则跳转至 3，否则跳转值 4
3. 直接插入 x
4. 将元素 2*x-min 插入到栈顶，更新 min=x

出栈操作：
设栈顶元素为 top，返回值为 x
1. 如果栈为空，抛出空栈异常
2. 如果 top >= min 直接 x=top
3. 如果 top < min，证明当前栈中的最小元素将要出栈，x=min，重新计算栈中的最小元素 min=x*2-y

注意入栈出栈的操作中，2*x-min 与 min=x*2-y 对称

工作原理:
如果插入元素 x < min
2*x-min < x 永远成立，(min > x)
而 newMin=x，则 newMin >= 2*x 永远成立

如果插入元素 x >= min
插入到栈顶的是 x，而 min <= x 永远成立

上述的 2 不能更改为 3 之类的数字

所以在 Pop 和 Peak 中，如果栈顶元素 < min，则可以判断需要更新 newMin = 2*min-top，出栈和进栈是一个对称操作
*/
type MinStack struct {
	elem []int
	min int
}

var EmptyStackError = errors.New("sorry MinStack is empty")

func (s *MinStack) Min() (int, error){
	if s.Len() == 0 {
		return 0, EmptyStackError
	}
	return s.min, nil
}

func (s *MinStack) Push(x int) {
	if s.Len() == 0 {
		s.elem = append(s.elem, x)
		s.min = x
		return
	}
	if x > s.min {
		s.elem = append(s.elem, x)
	} else {
		s.elem = append(s.elem, 2*x - s.min)
		s.min = x
	}
}

func (s *MinStack) Pop() (x int, err error) {
	if s.Len() == 0 {
		return 0, EmptyStackError
	}
	top := s.elem[len(s.elem)-1]
	s.elem = s.elem[:len(s.elem)-1]
	if top >= s.min {
		return top, nil
	}
	// 最小值被弹出栈，需要更新最小值
	x = s.min
	s.min = x * 2 - top
	return
}

func (s *MinStack) Peak() (x int, err error) {
	if s.Len() == 0 {
		return 0, EmptyStackError
	}
	top := s.elem[len(s.elem)-1]
	if top < s.min {
		x = s.min
	} else {
		x = top
	}
	return
}

func (s *MinStack) Len() int {
	return len(s.elem)
}