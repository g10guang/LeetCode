package sort

import (
	"testing"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	length := 100
	arr := make([]int, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range arr {
		arr[i] = r.Intn(100)
	}
	QuickSort(arr)
	//sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		assert.True(t, arr[i-1] <= arr[i])
	}
}

func TestQuickSort2(t *testing.T) {
	for i := 0; i < 10000; i++ {
		TestQuickSort(t)
	}
}
