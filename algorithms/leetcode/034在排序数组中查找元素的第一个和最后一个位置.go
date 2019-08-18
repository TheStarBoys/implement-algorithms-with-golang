package leetcode

import "math"
// 两种二分查找的实现， 第二种逻辑太复杂，容易出错
func searchRange(nums []int, target int) []int {
	first := search(nums, target)
	last := search(nums, target+1) - 1
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}else {
		return []int{first, int(math.Max(float64(first), float64(last)))}
	}
}

func search(nums []int, target int) int {
	l, h := 0, len(nums)
	for l < h {
		m := l + (h - l) / 2
		if nums[m] >= target {
			h = m
		}else {
			l = m + 1
		}
	}
	return l
}

func searchRange0(nums []int, target int) []int {
	r := []int{-1, -1}
	if len(nums) == 0 {
		return r
	}
	l, h := 0, len(nums) - 1
	for l < h {
		m := l + (h - l) / 2
		if nums[m] >= target {
			h = m
		}else {
			l = m + 1
		}
	}
	if nums[l] == target {
		if l >= len(nums) - 1 {
			return []int{l, l}
		}
		for i, v := range nums[l+1:] {
			if v != target {
				return []int{l, l + i}
			}
		}
		return []int{l, len(nums)-1}
	}
	return r
}
