package alg

import "errors"

type Stack interface {
	Len() int
	Top() (interface{}, error)
	IsEmpty() bool
	Push(interface{})
	Pop() (interface{}, error)
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) Top() (interface{}, error) {
	if s.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *stack) Pop() (interface{}, error) {
	if s.Len() == 0 {
		return nil, errors.New("empty stack")
	}
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x, nil
}
