package main

import "fmt"

type Tree struct {
	Val         int
	Left, Right *Tree
}

type QueueItem struct {
	Node  *Tree
	Level int
}

func width(t *Tree) int {
	if t == nil {
		return 0
	}
	var queue []*QueueItem
	numMap := make(map[int]int)
	queue = append(queue, &QueueItem{
		Node:  t,
		Level: 1,
	})
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		numMap[item.Level]++
		if item.Node.Left != nil {
			queue = append(queue, &QueueItem{
				Node:  item.Node.Left,
				Level: item.Level + 1,
			})
		}
		if item.Node.Right != nil {
			queue = append(queue, &QueueItem{
				Node:  item.Node.Right,
				Level: item.Level + 1,
			})
		}
	}
	maxVal := 0
	for _, v := range numMap {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	tree := &Tree{
		Left: &Tree{
			Left: &Tree{
				Val: 0,
				Left: &Tree{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &Tree{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Tree{
				Val: 0,
				Left: &Tree{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &Tree{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			Val: 1,
		},
		Right: &Tree{
			Left:  nil,
			Right: nil,
			Val:   9,
		},
		Val: 0,
	}
	fmt.Println(width(tree))
}

func test2() {
	fmt.Println(width(nil))
}

func test3() {
	tree := &Tree{
		Left: &Tree{
			Left: &Tree{
				Val: 0,
				Left: &Tree{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &Tree{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Tree{
				Val: 0,
				Right: &Tree{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			Val: 1,
		},
		Val: 0,
	}
	fmt.Println(width(tree))
}