package main

// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
type MinStack struct {
	container []*item
	minMark   []*item
}

type item struct {
	val int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	tm := &item{val: x}
	this.container = append(this.container, tm)
	if len(this.minMark) == 0 || this.minMark[len(this.minMark)-1].val >= tm.val {
		this.minMark = append(this.minMark, tm)
	}
}

func (this *MinStack) Pop() {
	tm := this.container[len(this.container)-1]
	this.container = this.container[:len(this.container)-1]
	min := this.minMark[len(this.minMark)-1]
	if tm == min {
		this.minMark = this.minMark[:len(this.minMark)-1]
	}
}

func (this *MinStack) Top() int {
	return this.container[len(this.container)-1].val
}

func (this *MinStack) Min() int {
	return this.minMark[len(this.minMark)-1].val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
