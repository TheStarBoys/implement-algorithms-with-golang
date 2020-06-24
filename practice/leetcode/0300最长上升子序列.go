package leetcode

// 回溯：超时
func lengthOfLIS0300_1(nums []int) int {
	var (
		ans int
		backtrack func(i, length int)
	)
	backtrack = func(i, length int) {
		for j := i+1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				backtrack(j, length+1)
			}
		}
		if length > ans {
			ans = length
		}
	}
	for i := 0; i < len(nums); i++ {
		backtrack(i, 1)
	}
	return ans
}

// 回溯：剪枝
func lengthOfLIS0300_2(nums []int) int {
	var (
		ans int
		backtrack func(i, length int)
		mem = make([]int, len(nums))
	)
	backtrack = func(i, length int) {
		if mem[i] > 0 && mem[i] > length {
			return
		}
		mem[i] = length
		if length > ans {
			ans = length
		}
		for j := i+1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				backtrack(j, length+1)
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		backtrack(i, 1)
	}
	return ans
}


// 动态规划
func lengthOfLIS0300_3(nums []int) int {
	if len(nums) == 0 { return 0 }
	dp := make([]int, len(nums))
	var ans int
	for i := 0; i < len(nums); i++ {
		tmp := 1
		// 从候选情况中找到最长的
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] { continue }
			tmp = max(tmp, dp[j]+1)
		}
		dp[i] = tmp
		ans = max(ans, tmp)
	}
	return ans
}