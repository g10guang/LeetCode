package sort_test

import (
	"testing"
	"github.com/g10guang/leetcode/sort"
	"strings"
)

func TestValidIpAddress(t *testing.T) {
	suits := []struct{
		input string
		expect []string
	} {
		{input:"25525511135", expect:[]string{"255.255.11.135", "255.255.111.35"}},
		{input:"0000", expect:[]string{"0.0.0.0"}},
		{input:"010010", expect:[]string{"0.10.0.10", "0.100.1.0"}},
	}
	for _, s := range suits {
		if r := sort.RestoreIpAddresses(s.input); !compareStringSlice(r, s.expect) {
			t.Errorf("input=%v expect=%v return=%v", s.input, s.expect, r)
		}
	}
}

func compareStringSlice(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if strings.Compare(s[i], t[i]) != 0 {
			return false
		}
	}
	return true
}
