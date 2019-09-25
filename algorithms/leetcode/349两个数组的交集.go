package leetcode

func intersection349_0(nums1 []int, nums2 []int) []int {
	set1, set2 := make(map[int]bool), make(map[int]bool)
	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}
	ans := []int{}
	if len(set1) <= len(set2) {
		for k := range set1 {
			if set2[k] {
				ans = append(ans, k)
			}
		}
	} else {
		for k := range set2 {
			if set1[k] {
				ans = append(ans, k)
			}
		}
	}
	return ans
}
