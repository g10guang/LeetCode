package sort_test

import (
	"testing"
	"github.com/g10guang/leetcode/sort"
)

func TestSortTopKFrequent(t *testing.T) {
	suits := []struct{
		input []int
		k int
		expect []int
	} {
		{[]int{1,1,1,2,2,3}, 2, []int{1,2}},
		{[]int{1,2}, 2, []int{1,2}},
	}
	for _, s := range suits {
		if r := sort.SortTopKFrequent(s.input, s.k); !compare(r, s.expect) {
			t.Errorf("input=%v k=%v expect=%v return=%v", s.input, s.k, s.expect, r)
		}
	}
}

func compare(s, t []int) bool {
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if v != t[i] {
			return false
		}
	}
	return true
}