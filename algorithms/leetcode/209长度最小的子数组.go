package leetcode

import "math"

// 效率很低，待优化
func minSubArrayLen209_0(s int, nums []int) int {
	sum, min := 0, math.MaxInt64
	slow := 0
	flag := 0
	for i := 0; i < len(nums); {
		if sum += nums[i]; sum >= s {
			if temp := i - slow + 1; min > temp {
				min = temp
			}
			sum = 0
			slow++
			i = slow
			flag = 1
			continue
		}
		i++
	}
	if flag == 0 {
		return 0
	}
	return min
}
