package alg_test

import (
	"testing"
	"github.com/g10guang/leetcode/alg"
)

func TestQuickSort(t *testing.T) {
	suits := generateData()
	for _, s := range suits {
		temp := make([]int, len(s.input))
		copy(temp, s.input)
		if alg.QuickSort(temp); !compare(temp, s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, temp)
		}
	}
}

func TestQuickSort2(t *testing.T) {
	suits := generateData()
	for _, s := range suits {
		temp := make([]int, len(s.input))
		copy(temp, s.input)
		if alg.QuickSort2(temp); !compare(temp, s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, temp)
		}
	}
}

func TestQuick3WaySort(t *testing.T) {
	suits := generateData()
	for _, s := range suits {
		temp := make([]int, len(s.input))
		copy(temp, s.input)
		if alg.Quick3WaySort(temp); !compare(temp, s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, temp)
		}
	}
}
