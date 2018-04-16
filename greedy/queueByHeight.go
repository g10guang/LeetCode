package greedy

import "sort"

type Int2DSlice [][]int

func (s Int2DSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Int2DSlice) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}
	return s[i][0] < s[j][0]
}

func (s Int2DSlice) Len() int {
	return len(s)
}

// https://leetcode.com/problems/queue-reconstruction-by-height/description/
func reconstructQueue(people [][]int) [][]int {
	sort.Sort(Int2DSlice(people))
	k := len(people) - 2
	for k >= 0 && people[k][0] == people[k+1][0] {
		k--
	}
	for k >= 0 {
		pos := people[k][1]
		t := k - 1
		for t >= 0 && people[t][0] == people[k][0] {
			pos--
			t--
		}
		tmp := people[k]
		for n, i := k, 0; pos > i; n, i = n+1, i+1 {
			people[k+i] = people[k+i+1]
		}
		people[k+pos] = tmp
		k--
	}
	return people
}

var ReconstructQueue = reconstructQueue