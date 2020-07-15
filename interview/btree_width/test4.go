package main

import "fmt"

type Tree struct {
	Val         int
	Left, Right *Tree
}

func width(root *Tree) int {
	if root == nil {
		return 0
	}
	var queue []*Tree
	maxWidth := 1
	queue = append(queue, root)
	for len(queue) > 0 {
		curWidth := len(queue)
		for curWidth > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			curWidth--
		}
		maxWidth = max(maxWidth, len(queue))
	}
	return maxWidth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
