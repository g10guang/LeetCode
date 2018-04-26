package alg_test

import (
	"testing"
	"github.com/g10guang/leetcode/alg"
)

func TestMinStack(t *testing.T) {
	s := alg.MinStack{}
	if s.Len() != 0 {
		t.Errorf("expect Len=0 not %d\n", s.Len())
	}
	s.Push(3)
	if s.Len() != 1 {
		t.Errorf("expect Len=1 not %d\n", s.Len())
	}
	r, err := s.Peak()
	if err != nil {
		t.Error(err)
	}
	if r != 3 {
		t.Errorf("Peak should be %d but %d", 10, r)
	}
	s.Push(5)
	r, err = s.Min()
	if err != nil {
		t.Error(err)
	}
	if r != 3 {
		t.Errorf("Min should be %d not %d", 3, r)
	}
	s.Push(2)
	s.Push(1)
	s.Push(1)
	s.Push(-1)
	r, err = s.Min()
	if err != nil {
		t.Error(err)
	}
	if r != -1 {
		t.Errorf("Min should be %d not %d", -1, r)
	}
	r = s.Len()
	if r != 6 {
		t.Errorf("Length should be %d not %d", 6, r)
	}
	r, err = s.Pop()
	if err != nil {
		t.Error(err)
	}
	if r != -1 {
		t.Errorf("Pop result should be %d not %d", -1, r)
	}
	r, err = s.Min()
	if err != nil {
		t.Error(err)
	}
	if r != 1 {
		t.Errorf("Min should be %d not %d", 1, r)
	}
	r, err = s.Pop()
	if err != nil {
		t.Error(err)
	}
	if r != 1 {
		t.Errorf("Pop result should be %d not %d", 1, r)
	}
	r, err = s.Pop()
	if err != nil {
		t.Error(err)
	}
	if r != 1 {
		t.Errorf("Pop result should be %d not %d", 1, r)
	}
	r, err = s.Pop()
	if err != nil {
		t.Error(err)
	}
	if r != 2 {
		t.Errorf("Pop result should be %d not %d", 2, r)
	}
	r, err = s.Pop()
	if err != nil {
		t.Error(err)
	}
	if r != 5 {
		t.Errorf("Pop result should be %d not %d", 5, r)
	}
	r = s.Len()
	if r != 0 {
		t.Errorf("Length should be %d not %d", 0, r)
	}
}
