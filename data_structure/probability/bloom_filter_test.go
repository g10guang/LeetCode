package probability

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBloomFilter(t *testing.T) {
	T := &bloomFilterStruct{}
	T.init(10)
	var set bloomFilter = T
	set.add("hello world")
	assert.True(t, set.isIn("hello world"))

	set.add("你好吗世界")
	assert.True(t, set.isIn("你好吗世界"))
}
