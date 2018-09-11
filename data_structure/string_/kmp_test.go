package string_

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKMP(t *testing.T) {
	assert.True(t, kmp("helloworld", "hello"))
	assert.True(t, kmp("BBC ABCDAB ABCDABCDABDE", "ABCDABD"))
	assert.True(t, kmp("ATCGATCGTTGA", "CGTTG"))
	assert.False(t, kmp("ATCGATCGTTGA", "CTG"))
}