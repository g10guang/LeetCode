package binarysearch

// https://leetcode.com/problems/arranging-coins/description/
// 抛硬币
func arrangeCoins(n int) int {
	l := 0
	h := n
	for l <= h {
		mid := l + (h-l)/2
		need := formula(mid)
		if need == n {
			return mid
		} else if need > n {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

func formula(x int) int {
	return (x + 1) * x / 2
}

var ArrageCoins = arrangeCoins
