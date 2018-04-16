package twoPointer_test

import (
	"testing"
	"github.com/g10guang/leetcode/twoPointer"
	"strings"
)

func TestReverseVowels(t *testing.T) {
	suits := []struct {
		input  string
		expect string
	}{
		{"EO", "OE"},
		{"hello", "holle"},
		{",.", ",."},
		{" ", " "},
		{"a.", "a."},
	}
	for _, s := range suits {
		if r := twoPointer.ReverseVowels(s.input); strings.Compare(r, s.expect) != 0 {
			t.Errorf("\n input=%v\n expect=%v\n return=%v\n", s.input, s.expect, r)
		}
	}
}
