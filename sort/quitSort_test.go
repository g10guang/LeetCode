package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid, i := data[0], 1
	head, tail := 0, len(data)-1
	for i = 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	qsort(data[:head])
	qsort(data[head+1:])
}

type S struct {
	input  []int
	expect []int
}

type T []int

func (t T) Len() int {
	return len(t)
}

func (t T) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t T) Less(i, j int) bool {
	return t[i] < t[j]
}

func generateRandomCases() []S {
	cases := make([]S, 10)
	for i := 0; i < 10; i++ {
		t := make([]int, 10)
		for j := 0; j < 10; j++ {
			t[j] = rand.Intn(100)
		}
		nt := make([]int, 10)
		copy(nt, t)
		sort.Sort(T(nt))
		cases = append(cases, S{t, nt})
	}
	return cases
}

func TestQuickSort(t *testing.T) {
	cases := generateRandomCases()
	for _, suit := range cases {
		qsort(suit.input)
		tmp := make([]int, 10)
		copy(tmp, suit.input)
		if !reflect.DeepEqual(suit.input, suit.expect) {
			t.Errorf("input=%v return=%v expected=%v\n", tmp, suit.input, suit.expect)
		} else {
			t.Logf("input=%v successful\n", tmp)
		}
	}
}
