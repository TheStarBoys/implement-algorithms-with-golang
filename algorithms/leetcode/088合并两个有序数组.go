package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	indexMerge := len(nums1) - 1
	for index1 >= 0 || index2 >= 0 {
		if index1 < 0 {
			nums1[indexMerge] = nums2[index2]
			indexMerge--
			index2--
		} else if index2 < 0 {
			nums1[indexMerge] = nums1[index1]
			indexMerge--
			index1--
		}else if nums1[index1] > nums2[index2] {
			nums1[indexMerge] = nums1[index1]
			indexMerge--
			index1--
		}else {
			nums1[indexMerge] = nums2[index2]
			indexMerge--
			index2--
		}
	}
}

func MergeTwoSortArrary(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}
