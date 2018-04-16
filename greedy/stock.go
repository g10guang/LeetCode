package greedy

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
// 这是最终答案
func maxProfit(prices []int) int {
	benefit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			benefit += prices[i] - prices[i-1]
		}
	}
	return benefit
}

func MaxProfit2(prices []int) int {
	benefit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			benefit += prices[i] - prices[i-1]
		}
	}
	return benefit
}

func MaxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	isCornor := findCornor(prices)
	isHold := false
	benefit := 0
	start := 0
	for i := 0; i < length-1; i++ {
		if isCornor[i] == 1 {
			// 卖出
			if isHold {
				isHold = false
				benefit += prices[i] - start
			}
		} else if isCornor[i] == -1 {
			// 买入
			if !isHold {
				isHold = true
				start = prices[i]
			}
		}
	}
	if isHold {
		benefit += prices[length-1] - start
	}
	return benefit
}

func findCornor(prices []int) []int {
	length := len(prices)
	// -1 谷 / 1 峰 / 0 两者都不是
	isCornor := make([]int, length)
	// 特殊对待第0个元素
	if prices[0] > prices[1] {
		isCornor[0] = 1
	} else if prices[0] < prices[1] {
		isCornor[0] = -1
	}
	// 特殊对待最后一个元素
	if prices[length-1] > prices[length-2] {
		isCornor[length-1] = 1
	} else if prices[length-1] < prices[length-2] {
		isCornor[length-1] = -1
	}
	// 计算其他元素
	for i := 1; i < length-1; i++ {
		switch {
		case prices[i] > prices[i+1] && prices[i] > prices[i-1]:
			isCornor[i] = 1
		case prices[i] < prices[i+1] && prices[i] < prices[i-1]:
			isCornor[i] = -1
		case prices[i] == prices[i-1]:
			if prices[i] > prices[i+1] {
				isCornor[i] = 1
			} else if prices[i] < prices[i+1] {
				isCornor[i] = -1
			}
		case prices[i] == prices[i+1]:
			if prices[i] > prices[i-1] {
				isCornor[i] = 1
			} else if prices[i] < prices[i-1] {
				isCornor[i] = -1
			}
		}
	}
	return isCornor
}
