package sort

import (
	"sort"
)

type charItem struct {
	letter rune
	val    int
}

type charSlice []*charItem

func (c *charSlice) Len() int {
	return len(*c)
}

func (c *charSlice) Less(i, j int) bool {
	// 倒序排序
	return (*c)[i].val > (*c)[j].val
}

func (c *charSlice) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *charSlice) insert(char rune) {
	(*c)[char].val++
}

func (c *charSlice) init() {
	for i := 0; i < len(*c); i++ {
		(*c)[i] = &charItem{letter: rune(i), val: 0}
	}
}

func frequencySort(s string) string {
	c := charSlice(make([]*charItem, 255))
	cs := &c
	cs.init()
	for _, v := range s {
		c.insert(v)
	}
	sort.Sort(cs)
	l := make([]rune, len(s))
	index := 0
	for _, v := range *cs {
		if v.val == 0 {
			break
		}
		for j := 0; j < v.val; j++ {
			l[index] = v.letter
			index++
		}
	}
	return string(l)
}

var FrequencySort = frequencySort
