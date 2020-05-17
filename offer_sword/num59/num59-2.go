package main

// https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/
type MaxQueue struct {
	queue []int		// normal container for FIFO
	deque []int		// mark which is the next biggest value
}


func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int)  {
	this.queue = append(this.queue, value)
	for idx := len(this.deque) - 1; idx >= 0 && this.deque[idx] < value; idx-- {
		this.deque = this.deque[:idx]
	}
	this.deque = append(this.deque, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	value := this.queue[0]
	this.queue = this.queue[1:]
	if value == this.Max_value() {
		this.deque = this.deque[1:]
	}
	return value
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */