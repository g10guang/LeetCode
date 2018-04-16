package greedy_test

import (
	"testing"
	"github.com/g10guang/leetcode/greedy"
)

func TestPartitionLabels(t *testing.T) {
	suits := []struct{
		input string
		expect []int
	} {
		{"ababcbacadefegdehijhklij", []int{9,7,8}},
	}
	for _, s := range suits {
		if r := greedy.PartitionLabels(s.input); !compareSlice(r, s.expect) {
			t.Errorf("\n input=%v\n expect=%v\n return=%v\n", s.input, s.expect, r)
		}
	}
}