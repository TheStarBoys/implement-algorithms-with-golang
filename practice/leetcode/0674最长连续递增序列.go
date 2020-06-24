package leetcode

// 滑动窗口
func findLengthOfLCIS0674_0(nums []int) int {
	var ans int
	l, r := 0, 1
	for ; r < len(nums); r++ {
		if nums[r-1] >= nums[r] {
			if ans < r - l {
				ans = r - l
			}
			l = r
		}
	}
	if r == len(nums) && l != r {
		if ans < r - l {
			ans = r - l
		}
	}
	return ans
}
