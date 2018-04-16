package alg_test

import (
	"testing"
	"github.com/g10guang/leetcode/alg"
	"time"
)

func TestConcurrency(t *testing.T) {
	q := alg.NoLockQueue{MaxSize:20}

	pchan := make(chan bool)
	cchan := make(chan bool)

	ps := make([]int, 0)
	cs := make([]int, 0)

	productor := func() {
		time.Sleep(time.Second)
		for i := 0; i < 200000; i++ {
			// 如果队列已经满了，不断尝试入队
			err := q.Enqueue(i)
			if err != nil {
				i--
			} else {
				ps = append(ps, i)
			}
		}
		pchan<- true
	}

	consumer := func() {
		time.Sleep(time.Second)
		for i := 0; i < 200000; i++ {
			x, err := q.Dequeue()
			if err != nil {
				i--
			} else {
				//t.Logf("%v\n", x)
				cs = append(cs, x.(int))
			}
		}
		cchan<-true
	}

	go productor()
	go consumer()

	pf, cf := true, true

	for pf || cf {
		select {
		case <-pchan:
			t.Logf("productor exit len=%v", q.Len())
			pf = false
		case <-cchan:
			t.Logf("consumer exit len=%v", q.Len())
			cf = false
		}
	}

	t.Logf("queue length=%v", q.Len())

	if !compare(ps, cs) {
		t.Errorf("进队出队序列不一致")
	}
}

func TestLinear(t *testing.T) {
	q := alg.NoLockQueue{MaxSize:200}
	for i := 0; i < 200; i++ {
		q.Enqueue(i)
	}
	if q.Len() != 200 {
		t.Errorf("Enqueue error want length=%v result=%v", 200, q.Len())
	}
	for i := 0; i < 100; i++ {
		x, err := q.Dequeue()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", x)
	}
	if q.Len() != 100 {
		t.Errorf("Dequeue error want length=%v result=%v", 100, q.Len())
	}
	for i := 0; i < 100; i++ {
		x, err := q.Dequeue()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", x)
	}
	if q.Len() != 0 {
		t.Errorf("Dequeue error want length=%v result=%v", 0, q.Len())
	}
}

