package priorityQueue

import (
	"testing"
	"strconv"
	"container/heap"
	"math/rand"
	"time"
)

func TestPriorityQueue_Len(t *testing.T) {
	q := PriorityQueue{}
	heap.Init(&q)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < 10; i++ {
		x := r.Intn(100)
		heap.Push(&q, &Item{Value: strconv.Itoa(x), Priority: x})
	}
	for q.Len() > 0 {
		item, ok := heap.Pop(&q).(*Item)
		if !ok {
			t.Fatalf("Type error\n")
		}
		t.Logf("item Value=%s Priority=%d\n", item.Value, item.Priority)
	}
}
