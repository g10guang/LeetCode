package tree

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func triesAdd(t *testing.T, T tries) {
	if T == nil {
		T = &trie{}
		T.init()
	}
	assert.True(t, T.add("hello"))
	assert.True(t, T.add("hel"))
	assert.True(t, T.add("world"))
	assert.False(t, T.add("hello"))
	assert.False(t, T.add("world"))
	assert.True(t, T.add("你好吗"))
}

func TestTriesAdd(t *testing.T) {
	triesAdd(t, nil)
}

func TestTriesIsIn(t *testing.T) {
	var T tries = &trie{}
	T.init()
	triesAdd(t, T)
	assert.True(t, T.isIn("hello"))
	assert.True(t, T.isIn("hel"))
	assert.True(t, T.isIn("world"))
	assert.False(t, T.isIn("hello world"))
	assert.False(t, T.isIn("world2"))
	assert.False(t, T.isIn("刘曦光"))
	assert.True(t, T.isIn("你好吗"))
}

func TestTriesRemove(t *testing.T) {
	var T tries = &trie{}
	T.init()
	triesAdd(t, T)
	assert.True(t, T.remove("hel"))
	assert.False(t, T.remove("world_"))
	assert.False(t, T.remove("worl"))
	assert.True(t, T.remove("hello"))
	assert.False(t, T.remove("helo"))
	assert.True(t, T.remove("world"))
	assert.False(t, T.remove("刘曦光"))
	assert.True(t, T.remove("你好吗"))
}

func TestTriesMaxPrefix(t *testing.T) {
	var T tries = &trie{}
	T.init()
	triesAdd(t, T)
	assert.True(t, T.maxPrefix("hello") == "hello")
	assert.True(t, T.maxPrefix("helloworld") == "hello")
	assert.True(t, T.maxPrefix("你好吗") == "你好吗")
	assert.True(t, T.maxPrefix("你好吗世界") == "你好吗")
	assert.True(t, T.maxPrefix("helo") == "hel")
	assert.True(t, T.maxPrefix("world") == "world")
	assert.True(t, T.maxPrefix("world?") == "world")
	assert.True(t, T.maxPrefix("no exists") == "")
	assert.True(t, T.maxPrefix("") == "")
}