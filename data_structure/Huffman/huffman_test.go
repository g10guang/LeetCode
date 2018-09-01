package Huffman

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

func TestTree_Build(t *testing.T) {
	ht := Tree{}
	length := 10
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	vals := make([]*Value, length)
	for i := 0; i < length; i++ {
		num := r.Intn(10)
		//vals = append(vals, &Value{Weight: num, Value: num})
		vals[i] = &Value{Weight: num, Value: num}
	}
	// print vals
	for _, v := range vals {
		fmt.Println(v)
	}
	ht.Init(vals)
	if !ht.Build() {
		t.Errorf("Build huffman tree faid\n")
	}
	traverse(t, &ht)
}

// traverse the huffman tree
func traverse(t *testing.T, ht *Tree) {
	curNode := ht.Root
	if curNode == nil {
		t.Errorf("Huffman tree root node is nil\n")
	}
	// breadth first traverse the huffman tree
	queue := make([]*Node, 1)
	queue[0] = curNode
	for len(queue) > 0 {
		tnode := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", tnode.Weight)
		if tnode.Left != nil {
			queue = append(queue, tnode.Left)
		}
		if tnode.Right != nil {
			queue = append(queue, tnode.Right)
		}
	}
	fmt.Println()
}
