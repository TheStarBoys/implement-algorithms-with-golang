package mysort

// 冒泡排序
// 时间O(N^2) 空间O(1)

// 内层循环每次两两比较，将最大值向右移到排序后的正确位置 len(nums) - i
// 由于内层循环每次将当前最大值放在最后，当外层循环到 len(nums) - 1 时，剩下最后一个元素，已经为正确位置
func bubbleSort(nums []int) {
	for i := 0; i < len(nums) - 1; i++ {
		for j := 1; j < len(nums) - i; j++ {
			if nums[j - 1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}