package greedy

// https://leetcode.com/problems/partition-labels/description/
func partitionLabels(S string) []int {
	// 因为题目说明只有 ascii 码，所以这里处理从简
	lastIndex := make([]int, 26)
	for i, c := range S {
		lastIndex[c-'a'] = i
	}
	partitions := make([]int, 0, 1)
	for j := 0; j < len(S); j++ {
		start := j
		end := lastIndex[S[j]-'a']
		for end > j {
			t := lastIndex[S[j]-'a']
			if t > end {
				end = t
			}
			j++
		}
		partitions = append(partitions, end-start+1)
		start = end + 1
	}
	return partitions
}

var PartitionLabels = partitionLabels

