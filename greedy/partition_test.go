package greedy_test

import (
	"testing"

	"github.com/g10guang/leetcode/greedy"
)

func TestPartition(t *testing.T) {
	suits := []struct {
		input  string
		expect []int
	}{
		 {"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"caedbdedda", []int{1, 9}},
	}
	for _, suit := range suits {
		r := greedy.PartitionLabels(suit.input)
		if !compareSlice(r, suit.expect) {
			t.Errorf("input=%v expect=%v result=%v", suit.input, suit.expect, r)
		}
	}
}

func compareSlice(s []int, t []int) bool {
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
