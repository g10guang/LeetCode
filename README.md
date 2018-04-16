本仓库记录 [LeetCode](https://leetcode.com/) 部分题目的解，其中使用的编程语言为 Go

其中题库跟随仓库：[https://github.com/CyC2018/Interview-Notebook/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3.md](https://github.com/CyC2018/Interview-Notebook/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3.md)

更加全面题解：[https://github.com/haoel/leetcode](https://github.com/haoel/leetcode)

# 目录

| [二分查找](#二分查找-binarySearch) | [贪心算法](#贪心算法-greedy) | [双指针](#双指针-twoPoiter) | [图的搜索](#图的遍历-graphSeach) | [回溯](#回溯-blackTrack) |
| :----- | :------- | :----- | :---- | :---- |


## 二分查找 :binarySearch:

### 求开方

[69. Sqrt(x)](https://leetcode.com/problems/sqrtx/description/)

[Solution](./binarysearch/sqrt.go)

### 抛硬币

[441. Arranging Coins](https://leetcode.com/problems/arranging-coins/description/)

[Soluton](./binarysearch/coins.go)

### 有序数组中的 Single Element

[540. Single Element in a Sorted Array](https://leetcode.com/problems/single-element-in-a-sorted-array/description/)

[Soluton](./binarysearch/singleNoDuplicate.go)

## 贪心算法 :greedy:

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

## 双指针 :twoPoiter:

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

## 排序 :sort:

### 第 K 个元素

[215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/description/)

[Solution](./sort/kth.go)

### Top-K Frequency elements

[347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/description/)

[Solution](./sort/topkFrequent.go)

### 根据字母频率排序

[451. Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency/description/)

[Solution](./sort/sortCharByFrequency.go)

## 图的遍历 :graphSeach:

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

## 回溯 :blackTrack:

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
