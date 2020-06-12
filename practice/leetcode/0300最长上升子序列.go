package leetcode

import "math"

// recursiveAlgorithm
func lengthOfLIS300_0(nums []int) int {
	return helper300_0(nums, -math.MaxInt64, 0)
}

func helper300_0(nums []int, prev, curIndx int) int {
	if curIndx >= len(nums) {
		return 0
	}
	tacken := 0
	if nums[curIndx] > prev {
		tacken = 1 + helper300_0(nums, nums[curIndx], curIndx+1)
	}
	notacken := helper300_0(nums, prev, curIndx+1)
	return int(math.Max(float64(tacken), float64(notacken)))
}

// 备忘录剪枝
func lengthOfLIS300_1(nums []int) int {
	// 行，保存prev的值；列，保存cur的值
	// len(nums)+1行， len(nums)列
	memo := make([][]int, len(nums)+1)
	tmp1 := make([]int, len(nums))
	for i := range tmp1 {
		tmp1[i] = -1
	}
	for i := range memo {
		tmp2 := make([]int, len(nums))
		copy(tmp2, tmp1)
		memo[i] = append(memo[i], tmp2...)
	}
	return dp300_1(nums, memo, -1, 0)
}

func dp300_1(nums []int, memo [][]int, prevIndx, curIndx int) int {
	if curIndx >= len(nums) {
		return 0
	}
	// 如果备忘录里有值
	if tmp := memo[prevIndx+1][curIndx]; tmp >= 0 {
		return tmp
	}
	tacken := 0
	if prevIndx < 0 || nums[curIndx] > nums[prevIndx] {
		tacken = 1 + dp300_1(nums, memo, curIndx, curIndx+1)
	}
	notacken := dp300_1(nums, memo, prevIndx, curIndx+1)
	// 将prevIndx对应的值右移一位，因为初始情况，prevIndx是-1,
	// 要让他能够在备忘录里不越界
	memo[prevIndx+1][curIndx] = int(math.Max(float64(tacken), float64(notacken)))
	return memo[prevIndx+1][curIndx]
}
