package dp

import (
	"fmt"
)

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

func multiKnapsack(w, p, n []int, N, W int) int {
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for v := 1; v <= W; v++ {
			for k := 0; k <= n[i]; k++ {
				if v >= k*w[i] {
					dp[i][v] = max(dp[i-1][v], dp[i-1][v-k*w[i]]+k*p[i])
				} else {
					dp[i][v] = dp[i-1][v]
				}
			}
		}
	}
	return dp[N][W]
}

func multiKnapsackSpaceAdvance(w, p, n []int, N, W int) int {
	dp := make([]int, W+1)
	for i := 0; i <= N; i++ {
		for v := W; v > 0; v-- {
			for k := 0; k <= n[i] && v >= k*w[i]; k++ {
				dp[v] = max(dp[v], dp[v-k*w[i]]+k*p[i])
			}
		}
	}
	return dp[W]
}

func multiDimensionKnapsack(w, p, s []int, N, W, S int) int {
	f := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		f[i] = make([][]int, W+1)
		for j := 0; j <= S; j++ {
			f[i][j] = make([]int, S+1)
		}
	}
	for i := 1; i <= N; i++ {
		for v := 0; v <= W; v++ {
			for y := 0; y <= S; y++ {
				if v >= w[i] && y >= s[i] {
					f[i][v][y] = max(f[i-1][v][y], f[i-1][v-w[i]][y-s[i]])
				} else {
					f[i][v][y] = f[i-1][v][y]
				}
			}
		}
	}
	return f[N][W][S]
}

func multiDimensionKnapsackSpaceAdvance(w, p, s []int, N, W, S int) int {
	f := make([][]int, W+1)
	for i := 0; i <= W; i++ {
		f[i] = make([]int, S+1)
	}
	for i := 1; i <= N; i++ {
		for v := W; v > 0 && v >= w[i]; v-- {
			for y := S; y > 0 && y >= s[i]; y-- {
				f[v][y] = max(f[v][y], f[v-w[i]][y-s[i]]+p[i])
			}
		}
	}
	return f[W][S]
}

func fullKnapsack(w []int, N, W int) bool {
	// dp mark every value is reachable or not.
	dp := make([]bool, W+1)
	// 0 can be reach always
	dp[0] = true
	// make sure every number will be used just once.
	for _, v := range w {
		// loop from W to 0
		// if loop from 0 to W, every dp[k * v] will be set true.
		for t := W; t > 0; t-- {
			if t >= v {
				dp[t] = dp[t] || dp[t-v]
			}
		}
	}
	// judge
	return dp[W]
}

func knapsackWaysCnt(w []int, N, W int) int {
	dp := make([]int, W+1)
	dp[0] = 1
	for _, v := range w {
		for t := W; t >= v; t-- {
			dp[t] += dp[t-v]
		}
	}
	return dp[W]
}

func comleteKnapsack(w, p []int, N, W int) int {
	dp := make([]int, W+1)
	for v := 1; v <= W; v++ {
		for i := range w {
			if v >= w[i] {
				dp[v] = max(dp[v], dp[v-w[i]]+p[i])
			}
		}
	}
	return dp[W]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
