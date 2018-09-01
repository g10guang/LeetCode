package twoPointer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	arr := []int{-1, 1, 2, 3, -4, 5, -6, 7}
	ret := ThreeSum(arr, 0	)
	assert.Equal(t, 4, ret)
}