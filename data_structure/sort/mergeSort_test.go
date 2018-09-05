package sort

import (
	"testing"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	length := 100
	arr := make([]int, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range arr {
		arr[i] = r.Intn(100)
	}
	MergeSort(arr)
	for i := 1; i < len(arr); i++ {
		assert.True(t, arr[i-1] <= arr[i])
	}
}

func TestMergeSort2(t *testing.T) {
	for i := 0; i < 1000; i++ {
		TestMergeSort(t)
	}
}