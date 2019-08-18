package leetcode

// 二分查找，算法思想比双指针复杂，但效率差不多
func singleNonDuplicate(nums []int) int {
	l, h := 0, len(nums) - 1
	for l < h {
		m := l + (h - l ) / 2;
		if m %2 == 1 {
			m-- // 保证 l/h/m 都在偶数位，使得查找区间大小一直都是奇数
		}
		if nums[m] == nums[m+1] {
			l = m + 2
		}else {
			h = m
		}
	}
	return nums[l]
}

// 双指针
func singleNonDuplicate1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}else if length == 1 {
		return nums[0]
	}
	l, r := 0, 1
	for r < length {
		if nums[l] != nums[r] {
			return nums[l]
		}
		l += 2
		r += 2
	}
	return nums[length-1]
}
// 双指针
func singleNonDuplicate2(nums []int) int {
	length := len(nums)
	l, r := 0, 1
	for r < length {
		if nums[l] != nums[r] {
			return nums[l]
		}
		l += 2
		r += 2
	}
	return nums[length-1]
}