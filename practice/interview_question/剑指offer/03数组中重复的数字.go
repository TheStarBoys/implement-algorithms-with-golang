package offer

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] { return v }
		m[v] = true
	}

	return -1
}
