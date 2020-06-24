package leetcode

import (
	"sort"
	"math"
)

// 排序+双指针
func threeSumClosest0016_0(nums []int, target int) int {
	// min(abs((a + b + c) - target))
	sort.Ints(nums)
	n := len(nums)
	ans := math.MaxInt32
	update := func(curr int) {
		if abs(curr - target) < abs(ans - target) {
			ans = curr
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] { continue }
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target { return target }
			update(sum)
			if sum < target {
				j++
				for ; j < k && nums[j] == nums[j-1]; j++ {}
			} else {
				k--
				for ; j < k && nums[k] == nums[k+1]; k-- {}
			}
		}
	}
	return ans
}
