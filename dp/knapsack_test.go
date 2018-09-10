package dp

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"time"
)

func zeroOneKnapsackInit() (weight []int, value []int, maxWeight, maxProfit int) {
	weight = []int{}
	value = []int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1000; i++ {
		weight = append(weight, r.Intn(1000))
		value = append(value, r.Intn(1000))
	}
	maxWeight = r.Intn(1000 * 5)
	maxProfit = 20
	return
}

func TestZeroOneKnapsack(t *testing.T) {
	weight, value, maxWeight, _ := zeroOneKnapsackInit()
	assert.Equal(t, zeroOneKnapsack(weight, value, maxWeight), zeroOneKnapsackSpaceAdvance(weight, value, maxWeight))
}
