package main 

import (
    "fmt"
)

// Q: https://leetcode-cn.com/problems/median-of-two-sorted-arrays/submissions/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) == 0 {
        return findGause(nums2)
    } else if len(nums2) == 0 {
        return findGause(nums1)
    }
    p1, p2 := 0, 0
    r1, r2 := len(nums1)-1, len(nums2)-1
    for p1 < r1 && p2 < r2 {
        if nums1[p1] <= nums2[p2] {
            p1++
        } else {
            p2++
        }

        if nums1[r1] >= nums2[r2] {
            r1--
        } else {
            r2--
        }
    }

    if p1 == r1 {
        return findLarange(nums1[p1], nums2[p2:r2+1])
    } else if p2 == r2 {
        return findLarange(nums2[p2], nums1[p1:r1+1])
    } else if p2 < r2 {
        return findGause(nums2[p2:r2+1])	
    } else {
        return findGause(nums1[p1:r1+1])
    }
}

func findLarange(n int, nums []int) float64 {
    p, r := 0, len(nums) - 1
    for p < r {
        if nums[p] <= n && nums[r] >= n {
            p++
            r--
        } else if nums[p] > n {
            return findGause(nums[p:r])
        } else if nums[r] < n {
            return findGause(nums[p+1:r+1])
        }
    }
    if p > r {
        return float64(n)
    }
    return float64(n + nums[p]) / 2
}

func findGause(nums []int) float64 {
    if len(nums) == 0 {
        return 0
    }
    p, r := 0, len(nums) - 1
    for p < r {
        p++
        r--
    }
    if p == r {
        return float64(nums[p])
    }
    return float64(nums[p] + nums[r]) / 2
}

func main() {
    fmt.Printf("result1=%v\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
    fmt.Printf("result2=%v\n", findMedianSortedArrays([]int{1, 3}, []int{2}))
    fmt.Printf("result3=%v\n", findMedianSortedArrays([]int{}, []int{2}))
    fmt.Printf("result4=%v\n", findMedianSortedArrays([]int{}, []int{}))
    fmt.Printf("result5=%v\n", findMedianSortedArrays([]int{1, 1, 1}, []int{}))
    fmt.Printf("result6=%v\n", findMedianSortedArrays([]int{1, 3, 5}, []int{4}))
    fmt.Printf("result7=%v\n", findMedianSortedArrays([]int{1, 2}, []int{-1, 3}))
    fmt.Printf("result8=%v\n", findMedianSortedArrays([]int{1, 1, 1}, []int{1, 1, 1}))
}