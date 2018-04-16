package blacktracking

import "strings"

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
func letterCombinations(digits string) []string {
	mapping := map[string][]string{
		"2": {"a", "b", "c"}, "3": {"d", "e", "f"},
		"4": {"g", "h", "i"}, "5": {"j", "k", "l"}, "6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"}, "8": {"t", "u", "v"}, "9": {"w", "x", "y", "z"},
	}
	nums := strings.Split(digits, "")
	length := len(nums)
	result := make([]string, 0)
	prefix := make([]string, length)
	prefix = prefix[:0]
	if length <= 0 {
		return result
	}
	var handle func(level int)
	handle = func(level int) {
		if level == length {
			result = append(result, strings.Join(prefix, ""))
			return
		}
		char := nums[level]
		letters, ok := mapping[char]
		if !ok {
			return
		}
		for _, s := range letters {
			prefix = append(prefix, s)
			handle(level+1)
			prefix = prefix[:len(prefix)-1]
		}
	}
	handle(0)
	return result
}

var LetterCombinations = letterCombinations
