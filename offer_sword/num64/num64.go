package main

import "fmt"

var result = 0
// https://leetcode-cn.com/problems/qiu-12n-lcof/
func sumNums(n int) int {
	result = 0
	sub(n)
	return result
}

func sub(n int) bool {
	_ = n > 0 && sub(n - 1)
	result += n
	return n > 0
}

func main() {
	fmt.Printf("%v\n", sumNums(9))
}