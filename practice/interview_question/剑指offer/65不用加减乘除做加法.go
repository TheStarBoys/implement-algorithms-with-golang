package offer

func add(a int, b int) int {
	if b == 0 {
		return a
	}
	// 非进位 + 进位
	return add(a ^ b, (a & b) << 1)
}
