package longest_common

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubstringInAString(t *testing.T) {
	assert.Equal(t, "ana", longestCommonSubtringInAString("banana"))
	assert.Equal(t, "l", longestCommonSubtringInAString("hello world"))
	assert.Equal(t, "e ", longestCommonSubtringInAString("how many students are in the classroom?"))
}

func TestCommonStringAtLeastMTimes(t *testing.T) {
	//assert.True(t, isEqual(map[string]interface{}{"a": nil, "ana": nil, "na": nil}, substringAtLeastMTimes("banana", 2)))
	//assert.True(t, isEqual(map[string]interface{}{"a": nil}, substringAtLeastMTimes("banana", 3)))
	//assert.True(t, isEqual(map[string]interface{}{"ab": nil, "b": nil}, substringAtLeastMTimes("ababab", 3)))
}

func TestLongestCommonSubstring(t *testing.T) {
	assert.Equal(t, "bc", longestCommonSubstring("abc", "bcf"))
	assert.Equal(t, "ana", longestCommonSubstring("banana", "Canada"))
	assert.Equal(t, "l", longestCommonSubstring("hello", "world"))
	assert.Equal(t, "", longestCommonSubstring("abc", "def"))
	assert.Equal(t, "", longestCommonSubstring("abc", ""))
}

func isEqual(expect map[string]interface{}, ret []string) bool {
	if len(expect) != len(ret) {
		return false
	}
	for _, v := range ret {
		if _, ok := expect[v]; !ok {
			return false
		}
	}
	return true
}
