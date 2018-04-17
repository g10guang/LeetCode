本仓库记录 [LeetCode](https://leetcode.com/) 部分题目的解，其中使用的编程语言为 Go

其中题库跟随仓库：[https://github.com/CyC2018/Interview-Notebook/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3.md](https://github.com/CyC2018/Interview-Notebook/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3.md)

更加全面题解：[https://github.com/haoel/leetcode](https://github.com/haoel/leetcode)

# 目录

- [目录](#%E7%9B%AE%E5%BD%95)
    - [二分查找](#%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE)
        - [求开方](#%E6%B1%82%E5%BC%80%E6%96%B9)
        - [抛硬币](#%E6%8A%9B%E7%A1%AC%E5%B8%81)
        - [有序数组中的 Single Element](#%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E4%B8%AD%E7%9A%84-single-element)
    - [贪心算法](#%E8%B4%AA%E5%BF%83%E7%AE%97%E6%B3%95)
        - [饼干分配](#%E9%A5%BC%E5%B9%B2%E5%88%86%E9%85%8D)
        - [投飞镖刺破气球](#%E6%8A%95%E9%A3%9E%E9%95%96%E5%88%BA%E7%A0%B4%E6%B0%94%E7%90%83)
        - [股票收益最大](#%E8%82%A1%E7%A5%A8%E6%94%B6%E7%9B%8A%E6%9C%80%E5%A4%A7)
        - [种植花朵](#%E7%A7%8D%E6%A4%8D%E8%8A%B1%E6%9C%B5)
        - [修改数组为非递减](#%E4%BF%AE%E6%94%B9%E6%95%B0%E7%BB%84%E4%B8%BA%E9%9D%9E%E9%80%92%E5%87%8F)
        - [子串问题](#%E5%AD%90%E4%B8%B2%E9%97%AE%E9%A2%98)
        - [分割字符串](#%E5%88%86%E5%89%B2%E5%AD%97%E7%AC%A6%E4%B8%B2)
        - [根据身高排队](#%E6%A0%B9%E6%8D%AE%E8%BA%AB%E9%AB%98%E6%8E%92%E9%98%9F)
    - [双指针](#%E5%8F%8C%E6%8C%87%E9%92%88)
        - [有序数组中寻找两个数字等于特定数字](#%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E4%B8%AD%E5%AF%BB%E6%89%BE%E4%B8%A4%E4%B8%AA%E6%95%B0%E5%AD%97%E7%AD%89%E4%BA%8E%E7%89%B9%E5%AE%9A%E6%95%B0%E5%AD%97)
        - [旋转元音字母](#%E6%97%8B%E8%BD%AC%E5%85%83%E9%9F%B3%E5%AD%97%E6%AF%8D)
        - [寻找两数平方和等于第三个数](#%E5%AF%BB%E6%89%BE%E4%B8%A4%E6%95%B0%E5%B9%B3%E6%96%B9%E5%92%8C%E7%AD%89%E4%BA%8E%E7%AC%AC%E4%B8%89%E4%B8%AA%E6%95%B0)
        - [字符串回文](#%E5%AD%97%E7%AC%A6%E4%B8%B2%E5%9B%9E%E6%96%87)
        - [归并排序数组](#%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F%E6%95%B0%E7%BB%84)
        - [链表是否有环](#%E9%93%BE%E8%A1%A8%E6%98%AF%E5%90%A6%E6%9C%89%E7%8E%AF)
        - [字符串最长单词](#%E5%AD%97%E7%AC%A6%E4%B8%B2%E6%9C%80%E9%95%BF%E5%8D%95%E8%AF%8D)
    - [排序](#%E6%8E%92%E5%BA%8F)
        - [第 K 个元素](#%E7%AC%AC-k-%E4%B8%AA%E5%85%83%E7%B4%A0)
        - [Top-K Frequency elements](#top-k-frequency-elements)
        - [根据字母频率排序](#%E6%A0%B9%E6%8D%AE%E5%AD%97%E6%AF%8D%E9%A2%91%E7%8E%87%E6%8E%92%E5%BA%8F)
    - [图的遍历](#%E5%9B%BE%E7%9A%84%E9%81%8D%E5%8E%86)
        - [岛屿的最大面积](#%E5%B2%9B%E5%B1%BF%E7%9A%84%E6%9C%80%E5%A4%A7%E9%9D%A2%E7%A7%AF)
        - [朋友圈](#%E6%9C%8B%E5%8F%8B%E5%9C%88)
        - [岛屿数量](#%E5%B2%9B%E5%B1%BF%E6%95%B0%E9%87%8F)
        - [二叉树叶子到根的路径](#%E4%BA%8C%E5%8F%89%E6%A0%91%E5%8F%B6%E5%AD%90%E5%88%B0%E6%A0%B9%E7%9A%84%E8%B7%AF%E5%BE%84)
        - [划分 IP 地址](#%E5%88%92%E5%88%86-ip-%E5%9C%B0%E5%9D%80)
        - [封闭区域](#%E5%B0%81%E9%97%AD%E5%8C%BA%E5%9F%9F)
        - [可达性分析](#%E5%8F%AF%E8%BE%BE%E6%80%A7%E5%88%86%E6%9E%90)
    - [回溯](#%E5%9B%9E%E6%BA%AF)
        - [数字组合](#%E6%95%B0%E5%AD%97%E7%BB%84%E5%90%88)
        - [单词搜索](#%E5%8D%95%E8%AF%8D%E6%90%9C%E7%B4%A2)
        - [排列组合](#%E6%8E%92%E5%88%97%E7%BB%84%E5%90%88)
        - [有重复元素下的排列组合](#%E6%9C%89%E9%87%8D%E5%A4%8D%E5%85%83%E7%B4%A0%E4%B8%8B%E7%9A%84%E6%8E%92%E5%88%97%E7%BB%84%E5%90%88)
        - [数字组合](#%E6%95%B0%E5%AD%97%E7%BB%84%E5%90%88)
        - [找出集和中和等于特殊值的子集和](#%E6%89%BE%E5%87%BA%E9%9B%86%E5%92%8C%E4%B8%AD%E5%92%8C%E7%AD%89%E4%BA%8E%E7%89%B9%E6%AE%8A%E5%80%BC%E7%9A%84%E5%AD%90%E9%9B%86%E5%92%8C)
        - [含有相同元素的集和，找和为特殊值的子集和](#%E5%90%AB%E6%9C%89%E7%9B%B8%E5%90%8C%E5%85%83%E7%B4%A0%E7%9A%84%E9%9B%86%E5%92%8C%EF%BC%8C%E6%89%BE%E5%92%8C%E4%B8%BA%E7%89%B9%E6%AE%8A%E5%80%BC%E7%9A%84%E5%AD%90%E9%9B%86%E5%92%8C)
        - [集和中寻找和为特殊值的子集和](#%E9%9B%86%E5%92%8C%E4%B8%AD%E5%AF%BB%E6%89%BE%E5%92%8C%E4%B8%BA%E7%89%B9%E6%AE%8A%E5%80%BC%E7%9A%84%E5%AD%90%E9%9B%86%E5%92%8C)

## 二分查找

### 求开方

[69. Sqrt(x)](https://leetcode.com/problems/sqrtx/description/)

[Solution](./binarysearch/sqrt.go)

### 抛硬币

[441. Arranging Coins](https://leetcode.com/problems/arranging-coins/description/)

[Soluton](./binarysearch/coins.go)

### 有序数组中的 Single Element

[540. Single Element in a Sorted Array](https://leetcode.com/problems/single-element-in-a-sorted-array/description/)

[Soluton](./binarysearch/singleNoDuplicate.go)

## 贪心算法

### 饼干分配

[455. Assign Cookies](https://leetcode.com/problems/assign-cookies/description/)

[Solution](./greedy/cookie.go)

### 投飞镖刺破气球

[452. Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/)

[Solution](./greedy/balloon.go)

### 股票收益最大

[122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/)

[Solution](/home/g10guang/go/src/github.com/g10guang/leetcode/greedy/stock.go)

### 种植花朵

[605. Can Place Flowers](https://leetcode.com/problems/can-place-flowers/description/)

[Solution](./greedy/flower.go)

### 修改数组为非递减

[665. Non-decreasing Array](https://leetcode.com/problems/non-decreasing-array/description/)

[Solution](/home/g10guang/go/src/github.com/g10guang/leetcode/greedy/no-decrease.go)

### 子串问题

[392. Is Subsequence](https://leetcode.com/problems/is-subsequence/description/)

[Solution](./greedy/isSequence.go)

### 分割字符串

[763. Partition Labels](https://leetcode.com/problems/partition-labels/description/)

[Solution](./greedy/partitionLabel.go)

### 根据身高排队

[406. Queue Reconstruction by Height](https://leetcode.com/problems/queue-reconstruction-by-height/description/)

[Solution](./greedy/queueByHeight.go)

## 双指针

### 有序数组中寻找两个数字等于特定数字

[167. Two Sum II - Input array is sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/)

[Solution](./twoPointer/twoSum.go)

### 旋转元音字母

[345. Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/description/)

[Solution](./twoPointer/reverseVowels.go)

### 寻找两数平方和等于第三个数

[633. Sum of Square Numbers](https://leetcode.com/problems/sum-of-square-numbers/description/)

[Solution](./twoPointer/judgeSquareSum.go)

### 字符串回文

[680. Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/description/)

[Solution](./twoPointer/validPalindrome.go)

### 归并排序数组

[88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/description/)

[Solution](./twoPointer/mergeArray.go)

### 链表是否有环

[141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/description/)

[Solution](/home/g10guang/go/src/github.com/g10guang/leetcode/twoPointer/isCyclic.c)

### 字符串最长单词

[524. Longest Word in Dictionary through Deleting](https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/description/)

[Solution](./twoPointer/findLongestWord.go)

## 排序

### 第 K 个元素

[215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/description/)

[Solution](./sort/kth.go)

### Top-K Frequency elements

[347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/description/)

[Solution](./sort/topkFrequent.go)

### 根据字母频率排序

[451. Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/description/)

[Solution](./sort/sortCharByFrequency.go)

## 图的遍历

### 岛屿的最大面积

[695. Max Area of Island](https://leetcode.com/problems/max-area-of-island/description/)

[Solution](./graph/maxAreaOfIsland.go)

### 朋友圈

[547. Friend Circles](https://leetcode.com/problems/friend-circles/description/)

[Solution](./graph/friendCircle.go)

### 岛屿数量

[200. Number of Islands](https://leetcode.com/problems/number-of-islands/description/)

[Solution](./graph/numIslands.go)

### 二叉树叶子到根的路径

[257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/description/)

[Solution](./graph/binaryTreePaths.go)

### 划分 IP 地址

[93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/description/)

[Solution](./sort/validIpAddress.go)

### 封闭区域

[130. Surrounded Regions](https://leetcode.com/problems/surrounded-regions/description/)

[Solution](/home/g10guang/go/src/github.com/g10guang/leetcode/graph/surround.go)

### 可达性分析

[417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/description/)

[Solution](./graph/pacific.go)

## 回溯

### 数字组合

[17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/)

[Solution](./blacktracking/letterCombination.go)

### 单词搜索

[79. Word Search](https://leetcode.com/problems/word-search/description/)

[Solution](./blacktracking/wordSearch.go)

### 排列组合

[46. Permutations](https://leetcode.com/problems/permutations/description/)

[Solution](./blacktracking/permutation.go)

### 有重复元素下的排列组合

[47. Permutations II](https://leetcode.com/problems/permutations-ii/description/)

[Solution](./blacktracking/permutation2.go)

### 数字组合

[77. Combinations](https://leetcode.com/problems/combinations/description/)

[Solution](./blacktracking/combination.go)

### 找出集和中和等于特殊值的子集和

[39. Combination Sum](https://leetcode.com/problems/combination-sum/description/)

[Solution](./blacktracking/combinationSum.go)

### 含有相同元素的集和，找和为特殊值的子集和

[40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/description/)

[Solution](./blacktracking/combinationSum2.go)

### 集和中寻找和为特殊值的子集和

[216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/description/)

[Solution](./blacktracking/combination3.go)
