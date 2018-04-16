package alg_test

import "testing"

import (
	"math/rand"
	"sort"
	"github.com/g10guang/leetcode/alg"
)

type intSlice []int

func (s *intSlice) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *intSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *intSlice) Len() int {
	return len(*s)
}

type Data struct {
	input  []int
	expect []int
}

func compare(s, t []int) bool {
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

func random(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(100)
	}
	return s
}

func TestSelectSort(t *testing.T) {
	suits := generateData()
	for _, s := range suits {
		temp := make([]int, len(s.input))
		copy(temp, s.input)
		if alg.SelectSort(temp); !compare(temp, s.expect) {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, temp)
		}
	}
}

func generateData() []Data {
	suits := []Data{
		{input: random(20), expect: make([]int, 20)},
		{input: random(41), expect: make([]int, 41)},
		{input: random(0), expect: make([]int, 0)},
	}
	for i := 0; i < len(suits); i++ {
		copy(suits[i].expect, suits[i].input)
		t := intSlice(suits[i].expect)
		sort.Sort(&t)
		suits[i].expect = []int(t)
	}
	return suits
}
