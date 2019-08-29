package leetcode

func merge088_0(nums1 []int, m int, nums2 []int, n int) {
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
// 解法2：
// 解题思路：利用插入排序的思想，把nums1当成是手牌，nums2当成是牌堆
// 从牌堆里面取出一张牌，跟手牌从右往左进行比较，最终将牌插到手牌中应该排在的位置
func merge088_1(nums1 []int, m int, nums2 []int, n int)  {
	for _, key := range nums2 {
		j := m - 1
		for j >= 0 && nums1[j] > key {
			nums1[j + 1] = nums1[j]
			j--
		}
		nums1[j + 1] = key
		m++
	}
}

// 解法3：双指针/从前往后
func merge088_2(nums1 []int, m int, nums2 []int, n int)  {
	nums1_copy := make([]int, m)
	copy(nums1_copy, nums1)
	p := 0
	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		if nums1_copy[p1] < nums2[p2] {
			nums1[p] = nums1_copy[p1]
			p1++
		}else {
			nums1[p] = nums2[p2]
			p2++
		}
		p++
	}
	if p1 < m {
		copy(nums1[p:], nums1_copy[p1:])
	}
	if p2 < n {
		copy(nums1[p:], nums2[p2:])
	}
}