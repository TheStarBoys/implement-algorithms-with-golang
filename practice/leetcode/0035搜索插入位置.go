package leetcode
// 解法1：暴力遍历
func searchInsert035_0(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}else if target < nums[i] {
			return i
		}
	}
	return len(nums)
}
// 解法2：二分查找，思想较复杂容易出错
func searchInsert1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l+r)/2
		if target > nums[mid] { // 去右区间找
			l = mid + 1
		} else if target < nums[mid] { // 去左区间找
			r = mid - 1
		} else { // 找到了
			return mid
		}
	}
	return l
}
// 解法3：二分查找，套用模板，最后对l，h进行判断即可
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l, h := 0, len(nums) - 1
	for l + 1 < h {
		mid := l + (h - l) / 2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			h = mid
		}else {
			l = mid
		}
	}
	if nums[l] >= target { // 此时target在l的位置，或者target在l的左边，但应该插入到l的位置
		return l
	}else if target > nums[h] { // 此时target在h的右边
		return h + 1
	}else {	// 此时target == nums[h]或者target在l，h的中间，所以需要插在h处
		return h
	}
}