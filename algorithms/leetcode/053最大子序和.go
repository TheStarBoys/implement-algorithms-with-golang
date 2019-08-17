package leetcode

import "math"

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	preSum := nums[0]
	maxSum := preSum
	for i := 1; i < len(nums); i++ {
		if preSum > 0 {
			preSum += nums[i]
		}else { preSum = nums[i] }
		maxSum = int(math.Max(float64(preSum), float64(maxSum)))
	}
	return maxSum
}