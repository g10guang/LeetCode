package main

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if length == 0 || length == 2 {
		return true
	}
	root := postorder[length-1]
	index := length - 2
	for ; index >= 0; index-- {
		if postorder[index] < root {
			break
		}
	}
	if index >= 0 {
		for i := 0; i < index; i++ {
			if postorder[i] > root {
				return false
			}
		}
		if !verifyPostorder(postorder[:index+1]) {
			return false
		}
	}
	if index != length-2 {
		if !verifyPostorder(postorder[index+1 : length-1]) {
			return false
		}
	}
	return true
}
