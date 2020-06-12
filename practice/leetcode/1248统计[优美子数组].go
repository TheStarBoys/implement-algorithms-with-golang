package leetcode

// 暴力解
func numberOfSubarrays1248_0(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	hash := [2]int{}
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j >= 0; j-- {
			hash[nums[j]%2]++
			if hash[1] == k {
				res++
			}
		}
		hash = [2]int{}
	}

	return res
}

// 滑动窗口
func numberOfSubarrays1248_1(nums []int, k int) int {
	// l, r为左右指针
	l, r, res := 0, 0, 0
	oddCnt := 0 // 奇数个数
	for r < len(nums) {
		// 统计奇数个数
		if nums[r] % 2 == 1 {
			oddCnt++
		}
		r++
		// 如果奇数个数为k个
		if oddCnt == k {
			// 将 r 指针移动至下一个奇数的位置
			tmp := r
			for r < len(nums) && nums[r] % 2 == 0 {
				r++
			}
			// 右侧偶数个数 = 右侧当前奇数下标 - 右侧前一个奇数下标
			rightEvenCnt := r - tmp
			// 统计左侧偶数个数，左指针移动一次，偶数加1，直到碰见窗口中第一个奇数
			leftEvenCnt := 0
			for ; nums[l] % 2 == 0; l++ {
				leftEvenCnt++
			}
			// 第 1 个奇数左边的 leftEvenCnt 个偶数都可以作为优美子数组的起点
			// (因为第1个奇数左边可以1个偶数都不取，所以起点的选择有 leftEvenCnt + 1 种）
			// 第 k 个奇数右边的 rightEvenCnt 个偶数都可以作为优美子数组的终点
			// (因为第k个奇数右边可以1个偶数都不取，所以终点的选择有 rightEvenCnt + 1 种）
			// 所以该滑动窗口中，优美子数组左右起点的选择组合数为 (leftEvenCnt + 1) * (rightEvenCnt + 1)
			res += (leftEvenCnt + 1) * (rightEvenCnt + 1)
			// 此时 l 指向的是第 1 个奇数，因为该区间已经统计完了，因此 l 右移一位，oddCnt--
			l++
			oddCnt--
		}
	}

	return res
}

