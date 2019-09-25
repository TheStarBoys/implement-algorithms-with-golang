package mysort

// 冒泡排序
// 时间O(N) 空间O(1)
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		for j := 1; j < len(nums); j++ {
			if nums[j - 1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}