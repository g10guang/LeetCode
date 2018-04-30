package alg

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x != i {
			if nums[x] == x {
				return x
			}
			nums[x], nums[i] = x, nums[x]
			i--
		}
	}
	return 0
}

var FindDuplicate = findDuplicate

