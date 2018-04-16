package alg_test

import (
	"testing"
	"github.com/g10guang/leetcode/alg"
	"strings"
)

func TestTwoCircle(t *testing.T) {
	suits := []struct{
		source string
		expect string
	} {
		{"hello world hello", "hello"},
		{"", ""},
		{"abcdefg", ""},
	}
	for _, s := range suits {
		if r := alg.TwoCircle(s.source); strings.Compare(s.expect, r) != 0 {
			t.Errorf("source=%q\n expect=%q\n return=%q\n", s.source, s.expect, r)
		}
	}
}

func TestSuffixArray(t *testing.T) {
	suits := []struct{
		source string
		expect string
	} {
		{"hello world hello", "hello"},
		{"", ""},
		{"abcdefg", ""},
		{"bananana", "anana"},
	}
	for _, s := range suits {
		if r := alg.SuffixArray(s.source); strings.Compare(s.expect, r) != 0 {
			t.Errorf("source=%q\n expect=%q\n return=%q\n", s.source, s.expect, r)
		}
	}
}

func TestSubStringDuplicateMTimes(t *testing.T) {
	suits := []struct{
		source string
		M int
		expect string
	} {
		{"hello world hello", 2, "hello"},
		{"", 1,""},
		{"abcdefg", 1, "abcdefg"},
		{"bananana", 3, "ana"},
		{"bananana", 4, "a"},
	}
	for _, s := range suits {
		if r := alg.SubStringDuplicateMTimes(s.source, s.M); strings.Compare(s.expect, r) != 0 {
			t.Errorf("source=%q\n expect=%q\n return=%q\n", s.source, s.expect, r)
		}
	}
}