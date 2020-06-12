package mysort

// 选择排序
func selectionSort(nums []int) {
	// i == len(nums) - 1时，已经为最后一个元素，不需要排序
	for i := 0; i < len(nums) - 1; i++ {
		// 记录最小元素下标
		min := i
		for j := i + 1; j < len(nums); j++ {
			// 更新最小元素下标
			if nums[min] > nums[j] {
				min = j
			}
		}
		// if  i != min { 此处可不做判断
		nums[i], nums[min] = nums[min], nums[i]
	}
}
