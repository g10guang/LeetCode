package main

import "encoding/json"

type LinkList struct {
	Val  int
	Next *LinkList
}

func quickSort(begin, end *LinkList) *LinkList {
	if begin != end {
		pivot := partition(begin, end)
		quickSort(begin, pivot)
		quickSort(pivot.Next, end)
	}
	return begin
}

func partition(begin, end *LinkList) *LinkList {
	pivot := begin.Val
	p := begin
	q := begin.Next
	for q != end {
		if q.Val < pivot {
			p = p.Next
			p.Val, q.Val = q.Val, p.Val
		}
		q = q.Next
	}
	p.Val, begin.Val = begin.Val, p.Val
	return p
}

func main() {
	test1()
}

func test1() {
	list := &LinkList{
		Val: 939,
		Next: &LinkList{
			Val: 444,
			Next: &LinkList{
				Val: 99,
				Next: &LinkList{
					Val: 0,
					Next: &LinkList{
						Val: 50,
					},
				},
			},
		},
	}
	begin := quickSort(list, nil)
	b, _ := json.Marshal(begin)
	println(string(b))
}
