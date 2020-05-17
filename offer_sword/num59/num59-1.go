package main

// https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	if k == 0 || len(nums) < k {
		return result
	}
	queue := Constructor()
	for idx := 0; idx < k; idx++ {
		queue.Push_back(nums[idx])
	}
	for idx := k; idx < len(nums); idx++ {
		result = append(result, queue.Max_value())
		queue.Pop_front()
		queue.Push_back(nums[idx])
	}
	result = append(result, queue.Max_value())
	return result
}

type MaxQueue struct {
	queue []int // normal container for FIFO
	deque []int // mark which is the next biggest value
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

func (this *MaxQueue) Push_back(value int) {
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
