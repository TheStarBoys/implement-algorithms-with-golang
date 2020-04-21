package leetcode

import "sort"
// 解法1：先排序，后比较邻近元素
func containsDuplicate217_0(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
// 解法2：利用map作为哈希集
func containsDuplicate217_1(nums []int) bool {
	visited := make(map[int]bool)
	for i := range nums {
		if visited[nums[i]] {
			return true
		}
		visited[nums[i]] = true
	}
	return false
}

// 解法3：朴素线性查找 【超时】
func containsDuplicate217_2(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] == nums[i] {return true}
		}
	}
	return false
}