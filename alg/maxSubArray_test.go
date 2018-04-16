package alg_test

import (
	"testing"
	"github.com/g10guang/leetcode/alg"
)

func TestMaxSubArrayDivideAndConquer(t *testing.T) {
	suits := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-1, -2, -3, -4 , -5}, -1},
		{[]int{-1,2,-3,4,-5,6}, 6},
	}
	for _, s := range suits {
		if r := alg.MaxSubArrayDivideAndConquer(s.input); r != s.expect {
			t.Errorf("input=%v\n expect=%v return=%v\n", s.input, s.expect, r)
		}
	}
}

func TestMaxSubArrayLinear(t *testing.T) {
	suits := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3, 4, -10}, 10},
		{[]int{-1, -2, -3, -4 , -5}, -1},
		{[]int{-1,2,-3,4,-5,6}, 6},
	}
	for _, s := range suits {
		if r := alg.MaxSubArrayLinear(s.input); r != s.expect {
			t.Errorf("input=%v\n expect=%v return=%v\n", s.input, s.expect, r)
		}
	}
}
