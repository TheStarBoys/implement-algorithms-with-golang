package leetcode

// 荷兰国旗问题
func sortColors075_0(nums []int)  {
	// p0为0的最右边界， p2为2的最左边界
	p0, p2, curr := 0, len(nums)-1, 0
	for curr <= p2 { // 注意，有=的情况
		switch nums[curr] {
		case 0:
			swap(nums, curr, p0)
			p0++
			curr++ // 注意，此处curr也需要自增
		case 1:
			curr++
		case 2:
			swap075_0(nums, curr, p2)
			p2--
		}
	}
}

func swap075_0(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
