package leetcode

// 标签说是要用栈，但没用，这样也是最优解
// 时间100% 空间91.18%
func nextGreaterElement496_0(nums1 []int, nums2 []int) []int {
	res := []int{}
	table := make(map[int]int) // int --> index
	for i, v := range nums2 {
		table[v] = i
	}
flag:
	for _, v := range nums1 {
		for indx2 := table[v]+1; indx2 < len(nums2); indx2++ {
			if nums2[indx2] > v {
				res = append(res, nums2[indx2])
				continue flag
			}
		}
		res = append(res, -1)
	}
	return res
}