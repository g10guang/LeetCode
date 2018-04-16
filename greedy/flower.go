package greedy

// https://leetcode.com/problems/can-place-flowers/description/
func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	count := 0
	// 第 0 位和最后一位特殊对待
	if l >= 2 {
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			flowerbed[0] = 1
			count++
		}
		if flowerbed[l-1] == 0 && flowerbed[l-2] == 0 {
			flowerbed[l-1] = 1
			count++
		}
	} else {
		return n <= 1 && flowerbed[0] == 0 || n  == 0
	}
	for i := 1; i < l - 1; i++ {
		if flowerbed[i] == 0 {
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count++
				i++ // 下一个坑已经没有希望了
			}
		} else {
			i++
		}
	}
	return count >= n
}