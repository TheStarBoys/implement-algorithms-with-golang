package leetcode

func removeElement027_0(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}
