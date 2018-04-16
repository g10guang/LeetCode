package alg

import (
	"strings"
	"sort"
)

// 使用两个循环寻找字符串中的最长子串
func TwoCircle(source string) string {
	maxLen := 0
	maxStart := 0
	for i := 0; i < len(source)-1; i++ {
		for j := i + 1; j < len(source); j++ {
			if r := commonLen(i, j, source); r > maxLen {
				maxLen = r
				maxStart = i
			}
		}
	}
	return source[maxStart : maxStart+maxLen]
}

// 寻找子串 s 和 t 从头部开始相同部分长度
func commonLen(i, j int, s string) (cnt int) {
	if i > j {
		i, j = j, i
	}
	for j < len(s) && s[j] == s[i] {
		cnt++
		i++
		j++
	}
	return cnt
}

// 基于后缀数组，寻找字符串中的重复出现的最长子串
func SuffixArray(source string) string {
	// 因为 Go 中指针不可运算，所以这里我们存储的是字符的起始位置
	sarr := make([]int, len(source))
	for i := 0; i < len(sarr); i++ {
		sarr[i] = i
	}
	ssa := StringSuffixArray{sarr:sarr, source:source}
	sort.Sort(&ssa)
	// 对排序后的子字符串，每个字符串与右边字符串进行比较，寻找字符串中重复出现的最长子串
	maxStart, maxLen := 0, 0
	for i := 0; i < len(sarr) - 1; i++ {
		if r := commonLen(sarr[i], sarr[i+1], source); r > maxLen {
			maxLen = r
			maxStart = sarr[i]
		}
	}
	return source[maxStart:maxStart+maxLen]
}

type StringSuffixArray struct {
	sarr   []int
	source string
}

func (s *StringSuffixArray) Less(i, j int) bool {
	return strings.Compare(s.source[s.sarr[i]:], s.source[s.sarr[j]:]) < 0
}

func (s *StringSuffixArray) Swap(i, j int) {
	s.sarr[i], s.sarr[j] = s.sarr[j], s.sarr[i]
}

func (s *StringSuffixArray) Len() int {
	return len(s.sarr)
}

// 基于后缀数组，寻找重复出现至少 M 次的最长子串
func SubStringDuplicateMTimes(source string, M int) string {
	if M <= 1 {
		return source
	}
	sarr := make([]int, len(source))
	for i := 0; i < len(sarr); i++ {
		sarr[i] = i
	}
	ssa := StringSuffixArray{sarr:sarr, source:source}
	sort.Sort(&ssa)
	maxStart, maxLen := 0, 0
	// 从后往前遍历
	for i := len(sarr) - 1; i - M + 1 >= 0; i-- {
		if r := commonLen(sarr[i], sarr[i-M+1], source); r > maxLen {
			maxLen = r
			maxStart = sarr[i]
		}
	}
	return source[maxStart:maxStart+maxLen]
}
