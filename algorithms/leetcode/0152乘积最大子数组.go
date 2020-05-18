package leetcode

import "math"

func maxProduct0152_0(nums []int) int {
	// 给定以 i 为结尾的子数组，有一下标 j 满足
	// 0 <= j <= i < len(nums), 使得nums[j:i+1] 为最大连续子数组
	// 如果知道 nums[j:i+1] 的乘积，可以递推出
	// nums[j-1:i+1] 的乘积 = nums[j:i+1] 的乘积 * nums[j-1]
	// 如果 i 为 len(nums) - 1, 即待求答案
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		tmp := 1
		for j := i; j >= 0; j-- {
			tmp *= nums[j]
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

// LeetCode 官方题解
func maxProduct0152_1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx * nums[i], max(mn * nums[i], nums[i]))
		minF = min(mn * nums[i], min(mx * nums[i], nums[i]))
		ans = max(maxF, ans)
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}