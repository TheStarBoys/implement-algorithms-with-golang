package offer

func twoSum(nums []int, target int) []int {
	l, h := 0, len(nums) - 1
	for l < h {
		sum := nums[l] + nums[h]
		if sum == target {
			return []int{nums[l], nums[h]}
		} else if sum > target {
			h--
		} else {
			l++
		}
	}

	return nil
}
