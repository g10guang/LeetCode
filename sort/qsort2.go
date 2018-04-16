package sort

func QuickSort(data []int) {
	mid := data[0] // 这可以是随机的 data[random] 一个数
	head := 0
	for i, tail := 1, len(data)-1; i <= tail; {
		if data[i] > mid {
			// data[i] should be on the right side of mid
			data[i], data[tail] = data[tail], data[i]
			tail-- // 因为 data[tail] 的位置可以固定了
			// 不能够移动 i，因为 data[i] 刚刚被转换过来，还不确定其所在位置
		} else {
			// data[i] <= mid，data[i] 应该在 mid 的左边，而最后 data[head]=mid，最后 head 就是 mid 所在的位置
			data[i], data[head] = data[head], data[i]
			// data[i] 已经固定
			i++
			// mid 最后应该的位置 head 也应该相应的右移
			head++
		}
	}
	data[head] = mid
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}
