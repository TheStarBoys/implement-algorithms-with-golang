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
// 把一个数组的元素放到哈希表中，在另一个数组的里面去遍历，查看是否重复
func intersect350_1(nums1 []int, nums2 []int) []int {
	table := make(map[int]int) // num -- > count
	for _, v := range nums1 {
		table[v]++
	}
	res := []int{}
	for _, v := range nums2 {
		if table[v] > 0 {
			res = append(res, v)
			table[v]--
		}
	}
	return res
}
// 思路跟上面的一样，不过空间复杂度是较小的数组的
func intersect350_2(nums1 []int, nums2 []int) []int {
	nums := nums1
	if len(nums1) > len(nums2) { // find min
		nums = nums2
	}
	table := make(map[int]int) // num -- > count
	for _, v := range nums {
		table[v]++
	}
	nums = nums2
	if len(nums2) < len(nums1) { // find max
		nums= nums1
	}
	res := []int{}
	for _, v := range nums {
		if table[v] > 0 {
			res = append(res, v)
			table[v]--
		}
	}
	return res
}