package main

// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
type stack struct {
	container []int
}

func (this *stack) push(val int) {
	this.container = append(this.container, val)
}

func (this *stack) pop() int {
	length := this.length()
	if length == 0 {
		return -1
	}
	val := this.container[length-1]
	this.container = this.container[:length-1]
	return val
}

func (this *stack) length() int {
	return len(this.container)
}

type CQueue struct {
	input  stack
	output stack
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.input.push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.output.length() == 0 {
		this.transfer()
	}
	return this.output.pop()
}

// transfer ouput
func (this *CQueue) transfer() {
	for this.input.length() > 0 {
		this.output.push(this.input.pop())
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
