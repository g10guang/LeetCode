package alg_test

import (
	"testing"
	"github.com/g10guang/leetcode/alg"
)

func TestMergeSort(t *testing.T) {
	suits := generateData()
	for _, s := range suits {
		temp := make([]int, len(s.input))
		copy(temp, s.input)
		if alg.MergeSort(temp); !compare(temp, s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, temp)
		}
	}
}
