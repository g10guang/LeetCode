package binarysearch

// https://leetcode.com/problems/sqrtx/description/
// 求一个数的开方
// sqrt == x / sqrt
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l := 0
	h := x
	for l <= h {
		mid := l + (h-l)/2
		t := x / mid
		if t == mid {
			return mid
		} else if t < mid {
			// mid is too large
			h = mid - 1
		} else {
			// mid is too small
			l = mid + 1
		}
	}
	return l - 1
}

var MySqrt = mySqrt
