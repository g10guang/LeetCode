package main

import (
	"time"
)

// Q: https://leetcode-cn.com/problems/lfu-cache/submissions/

type pair struct {
	val        int
	key        int
	elem       *Node
	accessTime time.Time
	accessCnt  int
}

func (p *pair) Less(other Interface) bool {
	if o, ok := other.(*pair); ok {
		return p.accessCnt < o.accessCnt && p.accessTime.Before(o.accessTime)
	}
	return false
}

type LFUCache struct {
	cache     map[int]*pair
	frequency *BTree
	capacity  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:     make(map[int]*pair, capacity),
		frequency: New(),
		capacity:  capacity,
	}
}

func (cache *LFUCache) Get(key int) int {
	if p := cache.cache[key]; p == nil {
		return -1
	} else {
		cache.pushFront(p)
		return p.val
	}
}

func (cache *LFUCache) Put(key int, value int) {
	if p := cache.cache[key]; p != nil {
		p.val = value
		cache.insert(p)
	} else {
		cache.insert(&pair{
			key: key,
			val: value,
		})
	}
}

func (cache *LFUCache) insert(p *pair) {
	cache.evict()
	cache.cache[p.key] = p
	cache.pushBack(p)
}

func (cache *LFUCache) evict() {
	if len(cache.cache) == cache.capacity {
		elem := cache.frequency.Back()
		p := elem.Value.(*pair)
		cache.frequency.Remove(elem)
		delete(cache.cache, p.key)
		p.elem = nil
	}
}

func (cache *LFUCache) pushFront(p *pair) {
	p.accessTime = time.Now()
	p.accessCnt++
	if p.elem != nil {
		cache.frequency.Remove(p.elem)
	}
	p.elem = cache.frequency.PushFront(p)
}

func (cache *LFUCache) pushBack(p *pair) {
	if p.elem != nil {
		cache.frequency.Remove(p.elem)
	}
	p.elem = cache.frequency.PushBack(p)
}

// case1: ["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
//[[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]
// case2: ["LFUCache","put","put","put","put","get"]
//[[2],[3,1],[2,1],[2,2],[4,4],[2]]

type Interface interface {
	Less(other Interface) bool
}

// unbalanced btree
type BTree struct {
	*Node
}

func New() *BTree {
	return &BTree{}
}

func (tree *BTree) insert(v Interface) {
	if tree.Node == nil {
		tree.Node = &Node{
			parent: nil,
			right:  nil,
			left:   nil,
			Val:    v,
		}
		return
	}

	tree.Node.insert(v)
}

func (tree *BTree) remove(node *Node) {
	if tree.Node == node {
		tree.Node = nil
	}
}

type Node struct {
	parent *Node
	right  *Node
	left   *Node
	Val    Interface
}

func (node *Node) insert(v Interface) {
	if node.Val.Less(v) {
		if node.right != nil {
			node.right.insert(v)
		} else {
			node.right = &Node{
				parent: node,
				right:  nil,
				left:   nil,
				Val:    v,
			}
		}
	} else {
		if node.left != nil {
			node.left.insert(v)
		} else {
			node.left = &Node{
				parent: node,
				right:  nil,
				left:   nil,
				Val:    v,
			}
		}
	}
}

// remove itself from btree
func (node *Node) remove(target *Node) {
	if node.Val.Less(target.Val) {

	}
}
