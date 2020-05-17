package main

import (
	"fmt"
	"strconv"

	"sort"
)

type IntSlice struct {
	data []int
}

func (s *IntSlice) Less(i, j int) bool {
	a, b := int2Slice(s.data[i]), int2Slice(s.data[j])
	m, n := a, b
	for len(a) > 0 && len(b) > 0 {
		x := a[len(a)-1]
		y := b[len(b)-1]
		a = a[:len(a)-1]
		b = b[:len(b)-1]
		if x == y {
			continue
		} else if x > y {
			return false
		} else {
			return true
		}
	}
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) == 0 {
		return m[len(m)-1] < b[len(b)-1]
	}
	return a[len(a)-1] <= n[len(n)-1]
}

func (s *IntSlice) Len() int {
	return len(s.data)
}

func (s *IntSlice) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func int2Slice(n int) []int {
	if n == 0 {
		return []int{0}
	}
	var result []int
	for n > 0 {
		result = append(result, n%10)
		n /= 10
	}
	return result
}

// https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
func minNumber(nums []int) string {
	s := &IntSlice{
		data: nums,
	}
	sort.Sort(s)
	// if s.data[len(s.data)-1] == 0 {
	// 	return "0"
	// }
	// move 0 to back
	// last0Index := -1
	// for i, v := range s.data {
	// 	if v != 0 {
	// 		break
	// 	}
	// 	last0Index = i
	// }
	// if last0Index != -1 {
	// 	s.data[0], s.data[last0Index+1] = s.data[last0Index+1], s.data[0]
	// }
	// fmt.Printf("%+v\n", s.data)
	var result string
	for _, v := range s.data {
		result += strconv.Itoa(v)
	}
	return result
}

func main() {
	fmt.Printf("%s\n", minNumber([]int{10, 2}))
	fmt.Printf("%s\n", minNumber([]int{3,30,34,5,9}))
	fmt.Printf("%s\n", minNumber([]int{1,2,3,4,5,6,7,8,9,0}))
	fmt.Printf("%s\n", minNumber([]int{0, 0}))
	fmt.Printf("%s\n", minNumber([]int{824,938,1399,5607,6973,5703,9609,4398,8247}))
	fmt.Printf("%s\n", minNumber([]int{12,121}))
}