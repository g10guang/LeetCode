package sort

import (
	"testing"
	"math/rand"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := 100
	arr := make([]int, length)
	for i := range arr {
		arr[i] = r.Intn(100)
	}
	InsertSort(arr)
	for i := 1; i < len(arr); i++ {
		assert.True(t, arr[i-1] <= arr[i])
	}
}
