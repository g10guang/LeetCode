package backtracking_test

import (
	"testing"
	"strings"
	"github.com/g10guang/leetcode/backtracking"
)

func TestLetterCombinations(t *testing.T) {
	suits := []struct{
		input string
		expect []string
	}{
		{input:"23", expect:[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, s := range suits {
		if r := backtracking.LetterCombinations(s.input); !compareStringSlice(r, s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, r)
		}
	}
}

func compareStringSlice(s,t []string) bool {
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if strings.Compare(t[i], v) != 0 {
			return false
		}
	}
	return true
}
