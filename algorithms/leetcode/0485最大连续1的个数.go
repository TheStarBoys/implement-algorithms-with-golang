package leetcode

func findMaxConsecutiveOnes485_0(nums []int) int {
	max := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
			continue
		}
		count++
		if max < count {
			max = count
		}
	}
	return max
}
