package leetcode
import "sort"

// 对于使min(a, b)和最大，必有nums排序后，两个相邻元素配对后，最小元素相加
// 即a0 + a2 + a4 + ... + an-2的和最大
func arrayPairSum561_0(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}
	return sum
}
