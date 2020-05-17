package main

import "fmt"

// https://leetcode-cn.com/problems/chou-shu-lcof/
func nthUglyNumber(n int) int {
	slice := make([]int, n)
	p2, p3, p5 := 0, 0, 0
	slice[0] = 1
	for idx := 1; idx < n; idx++ {
		a, b, c := slice[p2]*2, slice[p3]*3, slice[p5]*5
		t := min(a, b, c)
		fmt.Printf("%v %v %v %v\n", a, b, c, t)
		slice[idx] = t
		if t == a {
			p2++
		}
		if t == b {
			p3++
		}
		if t == c {
			p5++
		}
	}
	return slice[n-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func main() {
	fmt.Printf("%v\n", nthUglyNumber(10))
	fmt.Printf("%v\n", nthUglyNumber(1))
}
