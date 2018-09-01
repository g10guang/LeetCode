package LRU

import (
	"container/list"
	"sync"
)

type node struct {
	k        string
	v        interface{}
	listElem *list.Element
}

type Cache struct {
	pool    map[string]*node
	order   list.List // mark down the visit order.
	maxSize uint      // max elem count in the cache.
	size    uint      // current elem count int the cache.
	mutex   sync.Mutex
}

func (o *Cache) Insert(k string, v interface{}) {
	newNode := node{
		k:        k,
		v:        v,
		listElem: nil,
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if val, ok := o.pool[k]; ok {
		//fmt.Printf("k=%s conflict\n", k)
		// remove the old cache obj
		o.order.Remove(val.listElem)
		o.size--
	}
	//fmt.Printf("k=%s insert\n", k)
	for o.size >= o.maxSize {
		// remove some elem in the lru cache
		o.removeLruElem()
	}
	o.pool[k] = &newNode
	newNode.listElem = o.order.PushFront(&newNode)
	o.size++
}

// Remove the least recent use element.
// Must hold the wLock before exec this function
func (o *Cache) removeLruElem() {
	e := o.order.Back()
	if e != nil {
		t, ok := e.Value.(*node)
		if !ok {
			// error
			panic("error type\n")
		}
		//fmt.Printf("remove the lru elem k=%s\n", t.k)
		delete(o.pool, t.k)
		o.order.Remove(e)
		o.size--
	}
}

func (o *Cache) Remove(k string) bool {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if val, ok := o.pool[k]; ok {
		o.order.Remove(val.listElem)
		delete(o.pool, k)
		o.size--
		//fmt.Printf("remove a elem k=%s\n", k)
		return true
	}
	return false
}

func (o *Cache) Update(k string, v interface{}) {
	o.Insert(k, v)
}

func (o *Cache) Get(k string) (v interface{}, ok bool) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if v, ok = o.pool[k]; ok {
		t := v.(*node)
		// update visit order
		o.order.MoveBefore(t.listElem, o.order.Front())
		var n *node
		n, ok = v.(*node)
		v = n.v
	}
	return
}

func (o *Cache) Init(maxSize uint) {
	o.pool = make(map[string]*node)
	o.maxSize = maxSize
	o.size = 0
}

// In multi-thread environment this func is not precious.
func (o *Cache) getSize() uint {
	return o.size
}
