package LRU

import (
	"testing"
	"strconv"
)

func handle(t *testing.T) {
	c := Cache{}
	c.Init(3)
	// insert
	for i := 0; i < 10; i++ {
		tmp := strconv.Itoa(i)
		c.Insert(tmp, tmp)
		t.Logf("size=%d\n", c.size)
	}
	// check order
	t.Logf("Order in Cache\n")
	for e, i := c.order.Front(), 0; e != nil; e, i = e.Next(), i + 1 {
		tmp := e.Value.(*node)
		t.Logf("%d k=%s size=%d\n", i, tmp.k, c.size)
	}
	// get
	t.Logf("Find in the map\n")
	for i := 0; i < 10; i++ {
		tmp := strconv.Itoa(i)
		v, ok := c.Get(tmp)
		if !ok {
			t.Logf("k=%s not in cache\n", tmp)
		}
		str, ok := v.(string)
		c.Remove(tmp)
		t.Logf("k=%s v=%s size=%d\n", tmp, str, c.size)
	}
}

func TestLRUCache_Insert_MultiThread(t *testing.T) {
	for i := 0; i < 100; i++ {
		go handle(t)
	}
}

func TestCache_Insert(t *testing.T) {
	handle(t)
}