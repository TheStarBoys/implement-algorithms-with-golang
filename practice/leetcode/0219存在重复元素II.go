package leetcode

import "math"

// 时间复杂度过高，不推荐
func containsNearbyDuplicate219_0(nums []int, k int) bool {
	table := make(map[int][]int) // num --> []indx
	for i, v := range nums {
		table[v] = append(table[v], i)
	}
	for _, v := range table {
		if len(v) <= 1 { continue } // 筛选掉不重复的
		for i := range v { // []int
			for j := i+1; j < len(v); j++ {
				if getAbs219_0(v[i] - v[j]) <= k {
					return true
				}
			}
		}
	}
	return false
}

func getAbs219_0(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 该方案利用i-j的绝对值小于等于k, 来约束查找的区间，但仍然时间复杂度过高
func containsNearbyDuplicate219_1(nums []int, k int) bool {
	for i := range nums {
		for j := i + 1; j <= i + k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
// 线性搜索，维护了一个K大小的滑动窗口，超时
func containsNearbyDuplicate2(nums []int, k int) bool {
	for i := range nums {
		for j := int(math.Max(float64(i-k), 0)); j < i; j++ {
			if nums[i] == nums[j] { return true }
		}
	}
	return false
}
// sliding window + hashSet解法, 最优解
// 关键在于维护的一个大小为k的滑动窗口
// 时间O(n) 空间O(min(n,k))
func containsNearbyDuplicate219_3(nums []int, k int) bool {
	set := make(map[int]bool)
	for i, v := range nums {
		if set[v] { return true }
		set[v] = true
		if len(set) > k {
			delete(set, nums[i - k]) // 删除最旧的元素
		}
	}
	return false
}