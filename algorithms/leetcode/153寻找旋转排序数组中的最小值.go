package leetcode

// 二分查找，效率跟普通的差不多
func findMin(nums []int) int {
	l, h := 0, len(nums) - 1
	for l < h {
		m := l + (h - l) / 2
		if nums[m] <= nums[h] {
			h = m
		}else {
			l = m + 1
		}
	}
	return nums[h]
}

// 解题思路：
// 因为是旋转数组，旋转后，从右往左是降序。从右往左遍历到不符合降序的临界点，既是最小值。
// 如果长度为1， 则找不到临界点，返回nums[0]即可
func findMin0(nums []int) int {
	if len(nums) == 0 { // 此行可以省去，如果len == 0， 例题会报错
		return -1
	}
	for i := len(nums)-1; i >0; i-- {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}
