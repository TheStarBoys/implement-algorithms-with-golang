package leetcode
// 解法1：暴力遍历
func singleNumber136_0(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i] == nums[j] {
				count++
				break
			}
		}
		if count == 0 {
			return nums[i]
		}
		count = 0
	}
	return -1
}

// 解法2：利用map
func singleNumber136_1(nums []int) int {
	table := make(map[int]int)
	for i := range nums {
		table[nums[i]]++
	}
	for i := range nums {
		v := table[nums[i]]
		if v == 1 {
			return nums[i]
		}
	}
	return -1
}