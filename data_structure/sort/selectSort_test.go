package sort

import (
	"testing"
	"math/rand"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestSelectSort(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	length := 100
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = r.Intn(100)
	}
	SelectSort(arr)
	for i := 1; i < len(arr); i++ {
		assert.True(t, arr[i-1] <= arr[i])
	}
}