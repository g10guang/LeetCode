package dp

import "fmt"

// 0-1 knapsack.
// weight and value must be int
// time O(n^2) space O(n^2)
func zeroOneKnapsack(weight, value []int, maxWeight int) int {
	cnt := len(weight)
	if cnt != len(value) {
		panic(fmt.Sprintf("len(weight)=%d len(value)=%d not equal", cnt, len(value)))
	}
	// add 1 to to make it benefit for calculate.
	// use array not slice. Array need not to be made before use.
	// go compiler will init the array to be zero.
	profit := make([][]int, cnt+1)
	for i := 0; i < cnt+1; i++ {
		profit[i] = make([]int, maxWeight+1)
	}

	for i := 1; i <= cnt; i++ {
		for v := 1; v <= maxWeight; v++ {
			if weight[i-1] > v {
				profit[i][v] = profit[i-1][v]
			} else {
				// Local optimal leads to global optimal.
				profit[i][v] = max(profit[i-1][v], profit[i-1][v-weight[i-1]]+value[i-1])
			}
		}
	}

	// return counter of object is cnt and weight is maxWeight value.
	return profit[cnt][maxWeight]
}

func zeroOneKnapsackSpaceAdvance(weight, value []int, maxWeight int) int {
	cnt := len(weight)
	if cnt != len(value) {
		panic(fmt.Sprintf("len(weight)=%d len(value)=%d not equal", cnt, len(value)))
	}
	profit := make([]int, maxWeight+1)
	// Traverse every product.
	for i := 0; i < cnt; i++ {
		for v := maxWeight; v > 0; v-- {
			// If i == cnt - 1, the last product just need to traverse one time.
			// because we just care about profit[maxWeight]
			if weight[i] <= v {
				profit[v] = max(profit[v-weight[i]]+value[i], profit[v])
			}
		}
	}
	return profit[maxWeight]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
