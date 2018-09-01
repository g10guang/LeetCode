package twoPointer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestZeroSum(t *testing.T) {
	arr := []int{9, 10, -9, 20, -10, 1, -1}
	cnt := ZeroSum(arr, 0)
	assert.Equal(t, 3, cnt)

	arr = []int{1, 1, 1, -1, -1}
	cnt = ZeroSum2(arr, 0)
	assert.Equal(t, 6, cnt)
}
