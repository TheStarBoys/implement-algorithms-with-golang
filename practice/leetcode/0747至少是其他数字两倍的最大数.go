package leetcode

func dominantIndex747_0(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	maxIndex := 0
	for i := range nums {
		if max < nums[i] {
			max = nums[i]
			maxIndex = i
		}
	}
	nums = append(nums[:maxIndex], nums[maxIndex + 1:]...)
	for i := range nums {
		if nums[i] * 2 > max {
			return -1
		}
	}
	return maxIndex
}
