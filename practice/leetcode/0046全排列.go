package leetcode

func permute046_0(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	return backtrack046_0(n, nums, ans, 0)
}

func backtrack046_0(n int, nums []int, ans [][]int, first int) [][]int{
	if first == n {
		tmp := make([]int, n)
		copy(tmp, nums)
		// 必须用一个临时的变量存值, 因为ans里的[]int元素实际是一个引用
		// 如果append(ans, nums), nums的值变化了，ans里的值也会跟着变化
		return append(ans, tmp)
	}
	for i := first; i < n; i++ {
		swap(nums, first, i)
		ans = backtrack046_0(n, nums, ans, first + 1)
		// 回溯
		swap(nums, first, i)
	}
	return ans
}

func swap046_0(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
