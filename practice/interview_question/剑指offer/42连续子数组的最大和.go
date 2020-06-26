package offer

import "math"

func maxSubArray(nums []int) int {
	currNum := 0
	ans := math.MinInt32
	for i := range nums {
		if currNum < 0 {
			currNum = nums[i]
		} else {
			currNum += nums[i]
		}
		ans = max(ans, currNum)
	}
	return ans
}
