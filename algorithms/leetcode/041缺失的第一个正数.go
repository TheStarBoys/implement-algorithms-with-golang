package leetcode

import "math"

// 哈希集, 时间O(n) 空间O(n)
func firstMissingPositive041_0(nums []int) int {
	set := make(map[int]bool)
	for i := range nums {
		if nums[i] <= 0 {
			continue
		}
		set[nums[i]] = true
	}
	for i := 1; i < math.MaxInt64; i++ {
		if !set[i] {
			return i
		}
	}
	return -1
}

// todo 最优解待学习