package leetcode

func findMedianSortedArrays0004_0(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	i, j := 0, 0
	cur := 0
	nums := make([]int, len(nums1) + len(nums2))
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			nums[cur] = nums1[i]
			i++
		} else {
			nums[cur] = nums2[j]
			j++
		}
		cur++
	}
	if i == len(nums1) {
		copy(nums[cur:], nums2[j:])
	}
	if j == len(nums2) {
		copy(nums[cur:], nums1[i:])
	}
	if len(nums) % 2 == 0 {
		right := len(nums) / 2
		return float64(nums[right] + nums[right-1]) / 2
	} else {
		return float64(nums[len(nums) / 2])
	}
}
