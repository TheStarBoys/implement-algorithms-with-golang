package leetcode

// 解法1：暴力求解
func twoSum001_0(nums []int, target int) []int {
	for i, _ := range nums {
		for j, _ := range nums {
			if i == j {
				continue
			}
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
// 解法2：暴力求解，但比解法1省去了不少遍历的次数
func twoSum001_1(nums []int, target int) []int {
	for i := range nums {
		tar := target - nums[i]
		for j := i + 1; j < len(nums); j ++ {
			if nums[j] == tar {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
// 解法3：利用map，空间换时间
func twoSum001_2(nums []int, target int) []int {
	table := make(map[int]int)
	for i := range nums {
		table[nums[i]] = i
	}
	for i := range nums {
		j, ok := table[target - nums[i]]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return []int{-1, -1}
}
// 解法4：一遍map
func twoSum001_3(nums []int, target int) []int {
	table := make(map[int]int)
	for i := range nums {
		j, ok := table[target - nums[i]] // 先判断，再赋值，否则可能会出现覆盖
		if ok { // 且不需要判断 i和j是否相等
			return []int{i, j}
		}
		table[nums[i]] = i
	}
	return []int{-1, -1}
}