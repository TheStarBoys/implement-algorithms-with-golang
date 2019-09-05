package leetcode

import "math"
// 两种二分查找的实现， 第二种逻辑太复杂，容易出错
func searchRange034_0(nums []int, target int) []int {
	first := search(nums, target)
	last := search(nums, target+1) - 1
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}else {
		return []int{first, int(math.Max(float64(first), float64(last)))}
	}
}
// 查找到target的第一个下标
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
// 先二分查找到出现的第一个坐标位置， 再循环遍历找到最后一个位置
func searchRange034_1(nums []int, target int) []int {
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

// 同样是二分查找，不过此次是通过两次二分查找，分别插第一次的位置，跟最后一次的位置
func searchRange034_2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, h := 0, len(nums) - 1
	res := []int{}
	for l + 1 < h {
		mid := l + (h - l) / 2
		if nums[mid] >= target {
			h = mid
		}else {
			l = mid
		}
	}
	if nums[l] == target {
		res = append(res, l)
	}else if nums[h] == target {
		res = append(res, h)
	}else {
		return []int{-1, -1}
	}
	l, h = 0, len(nums) - 1
	for l + 1 < h {
		mid := l + (h - l) / 2
		if nums[mid] <= target {
			l = mid
		}else {
			h = mid
		}
	}
	if nums[h] == target {
		res = append(res, h)
	}else if nums[l] == target {
		res = append(res, l)
	}else {
		return []int{-1, -1}
	}
	return res
}
