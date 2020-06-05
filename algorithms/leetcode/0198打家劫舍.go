package leetcode

import "math"

// 动态规划
func rob0198_0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0] // 如果只有一间房，只能偷它
	dp[1] = max(nums[0], nums[1]) // 如果有两间房，偷两间房中较大金额的
	for i := 2; i < len(nums); i++ {
		// 对于房间 i , 如果偷房间 i 那么 i - 1 处不能偷, 那么所偷金额为 max(dp[i-2] + nums[i])
		// 对于房间 i , 如果不偷房间 i 那么所偷金额为 dp[i-1]
		// 对于房间 i , 偷这两种情况中的较大金额的情况
		dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// 当前偷的金额选择，只跟前两项有关
func rob0198_1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first + nums[i], second)
	}
	return second
}

// 删减对边界条件的判断的写法
func rob0198_2(nums []int) int {
	pre2, pre1 := 0, 0
	for i := 0; i < len(nums); i++ {
		cur := int(math.Max(float64(pre2 + nums[i]), float64(pre1)))
		pre2, pre1 = pre1, cur
	}
	return pre1
}
