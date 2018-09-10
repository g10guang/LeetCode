package longest_common

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	assert.Equal(t, 5, longestCommonSubsequence("hello world", "hello"))
	assert.Equal(t, 2, longestCommonSubsequence("你好吗世界", "好吗"))
	assert.Equal(t, 1, longestCommonSubsequence("你好吗世界", "不好"))
	assert.Equal(t, 12, longestCommonSubsequence("how are you my friend", "how are you lucy?"))
}
