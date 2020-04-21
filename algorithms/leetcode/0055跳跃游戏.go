package leetcode

// 递归
func canJump055_0(nums []int) bool {
	return dp055_0(nums, 0,len(nums)-1)
}

func dp055_0(nums []int, start, end int) bool {
	if start >= end { return true }
	for i := start+1; i <= nums[start] + start; i++ {
		if dp055_0(nums, i, end) {
			return true
		}
	}
	return false
}

// 带备忘录的递归
func canJump055_1(nums []int) bool {
	memo := make([]int, len(nums))
	// 能达到1, 不能达到-1, 未计算0
	return dp055_1(nums, memo, 0,len(nums)-1)
}

func dp055_1(nums, memo []int, start, end int) bool {
	if start >= end { return true }
	// 如果表没计算值，则计算
	if memo[start] == 0 {
		// 计算子问题，即当前位置所能跳到的范围之中，是否有能达到end的
		for i := start+1; i <= end && i <= nums[start] + start; i++ {
			if dp055_1(nums, memo, i, end) {
				memo[start] = 1
				break
			} else {
				memo[i] = -1
			}
		}
	}

	return memo[start] == 1
}

// 贪心
func canJump055_2(nums []int) bool {
	// 从右往左遍历，判断当前位置能不能跳到最后
	set := make(map[int]bool) // 能够跳到最后的下标集合
	for i := len(nums)-1; i >= 0; i-- {
		if nums[i] + i >= len(nums)-1 {
			set[i] = true
			continue
		}
		// 遍历i所能跳到的所有位置，有没有能跳到set集合已有的元素的
		j := nums[i] + i
		if j >= len(nums) {
			j = len(nums)-1
		}
		for ; j >= i; j-- {
			if set[j] {
				set[i] = true
				break
			}
		}
	}
	return set[0]
}
// 贪心，优化
// 实际上不需要集合来存储
// 如果跳跃的区间包含了minIndx，自然而然的能跳到最后
func canJump055_3(nums []int) bool {
	// 从右往左遍历，判断当前位置能不能跳到最后
	minIndx := len(nums)-1 // 能够跳到最后的最小下标
	for i := len(nums)-2; i >= 0; i-- {
		if nums[i] + i >= minIndx {
			minIndx = i
			continue
		}
	}
	return minIndx == 0
}
