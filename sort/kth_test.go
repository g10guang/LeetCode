package sort_test

import (
	"testing"
	"github.com/g10guang/leetcode/sort"
)

func TestHeapFindKthLargest(t *testing.T) {
	suits := []struct{
		input []int
		k int
		expect int
	} {
		{[]int{3,2,1,5,6,4}, 2, 5},
		{[]int{2,1}, 2, 1},
	}
	for _, s := range suits {
		if r := sort.HeapFindKthLargest(s.input, s.k); r != s.expect {
			t.Errorf("input=%v k=%v expect=%v return=%v\n", s.input, s.k, s.expect, r)
		}
	}
}


func TestQsortFindKthLargest(t *testing.T) {
	suits := []struct{
		input []int
		k int
		expect int
	} {
		{[]int{3,2,1,5,6,4}, 2, 5},
		{[]int{2,1}, 2, 1},
	}
	for _, s := range suits {
		if r := sort.QsortFindKthLargest(s.input, s.k); r != s.expect {
			t.Errorf("input=%v k=%v expect=%v return=%v\n", s.input, s.k, s.expect, r)
		}
	}
}