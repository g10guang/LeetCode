package alg_test

import (
	"testing"
	"github.com/g10guang/leetcode/alg"
)

func TestHeap_Sort(t *testing.T) {
	suits := generateData()
	for _, s := range suits {
		temp := make([]int, len(s.input))
		copy(temp, s.input)
		h := new(alg.Heap)
		h.MaxSize = len(temp)
		for _, v := range temp {
			h.Insert(v)
		}
		//h.BottomUpConstruct()
		h.UpBottomConstruct()
		h.Sort()
		if !compare(h.Elems(), s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, h.Elems())
		}
	}
}
