package offer

func majorityElement(nums []int) int {
	m := make(map[int]int)
	half := len(nums) >> 1
	for _, v := range nums {
		m[v]++
		if m[v] > half {
			return v
		}
	}

	return -1
}
