package twoPointer

import "math"

func judgeSquareSum(c int) bool {
	sqrt := int(math.Sqrt(float64(c)))
	p, q := 0, sqrt
	for p >= 0 && q >= 0 {
		t := p * p + q * q
		if t == c {
			return true
		} else if t < c {
			p++
		} else {
			q--
		}
	}
	return false
}
