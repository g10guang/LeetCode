package twoPointer_test

import (
	"testing"
	"github.com/g10guang/leetcode/twoPointer"
)

func TestMerge(t *testing.T) {
	suits := []struct{
		nums1 []int
		m int
		nums2 []int
		n int
		expect []int
	} {
		{nums1: []int{1, 3, 5, 7}, m: 4, nums2: []int{2, 4, 6}, n: 3, expect: []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, s := range suits {
		if r := twoPointer.Merge(s.nums1, s.m, s.nums2, s.n); !isEqual(r, s.expect) {
			t.Errorf("input: %v %v %v %v return: %v expect: %v", s.nums1, s.m, s.nums2, s.n, r, s.expect)
		}
	}
}

func isEqual(m, n []int) bool {
	if len(m) != len(n) {
		return false
	}
	for i, v := range m {
		if n[i] != v {
			return false
		}
	}
	return true
}
