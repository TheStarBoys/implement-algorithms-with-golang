package leetcode
// 解法1：每找到一个重复元素，就删除掉切片对应的元素
func intersect350_0(nums1 []int, nums2 []int) []int {
	dup := make([]int, 0)
	for i := 0; i < len(nums1); {
		for j := 0; j < len(nums2); {
			if nums1[i] == nums2[j] {
				dup = append(dup, nums1[i])
				nums1 = append(nums1[:i], nums1[i+1:]...)
				nums2 = append(nums2[:j], nums2[j+1:]...)
				i--
				break
			}
			j++
		}
		i++
	}
	return dup
}
