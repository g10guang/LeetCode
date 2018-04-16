package alg

import "math"

// 使用分治法求最大子数组问题
// 思路：最大子数组只可能存在三种情况：
// 1. 存在与左边的子数组中
// 2. 存在与右边的子数组中
// 3. 跨越左右两边
// 时间复杂度为：O(nlogn)
// 空间复杂度为：O(logn)
func MaxSubArrayDivideAndConquer(data []int) int {
	if len(data) == 0 {
		return math.MinInt32
	}
	mid := len(data) / 2
	// 计算跨越中介的最大子数组
	midmax, sum := math.MinInt32, 0
	for i := mid; i >= 0; i-- {
		sum += data[i]
		if sum > midmax {
			midmax = sum
		}
	}
	for j := mid + 1; j < len(data); j++ {
		sum += data[j]
		if sum > midmax {
			midmax = sum
		}
	}
	// 计算左数组中的最大值
	l := MaxSubArrayDivideAndConquer(data[:mid])
	// 计算右数组的最大值
	r := MaxSubArrayDivideAndConquer(data[mid+1:])
	var max int
	if l > r {
		max = l
	} else {
		max = r
	}
	if midmax > max {
		max = midmax
	}
	return max
}


/*
使用线性方法寻找数组中的最大子数组
思路：
最大子数组有可能以任何一个下标为结束 data[x:y]
*/
func MaxSubArrayLinear(data []int) int {
	max, leftSum := math.MinInt32, math.MinInt32
	for i := 0; i < len(data); i++ {
		if leftSum <= 0 {
			leftSum = 0
		}
		leftSum += data[i]
		if leftSum > max {
			max = leftSum
		}
	}
	return max
}
