package offer

func exchange(nums []int) []int {
	if len(nums) == 0 { return nil }
	n := len(nums)
	l, r := 0, n - 1
	for {
		for ; l < n && nums[l] & 1 == 1; l++ {}
		for ; r >= 0 && nums[r] & 1 == 0; r-- {}
		if l >= r {
			break
		}
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return nums
}
