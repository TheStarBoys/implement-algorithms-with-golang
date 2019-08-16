package leetcode

func removeDuplicates(nums []int) int {
	pre, cur := 0, 0
	if len(nums) == 0 {
		return 0
	}
	for cur < len(nums) {
		if nums[pre] == nums[cur] {
			cur++
		}else {
			pre++
			nums[pre] = nums[cur]
			cur++
		}
	}

	return pre + 1
}
