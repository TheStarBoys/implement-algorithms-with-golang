package offer

// 数学+递归。属于一看就会，一做就废的题，面试遇到还是弃了吧
func lastRemaining(n int, m int) int {
	if n == 1 { return 0 }
	x := lastRemaining(n-1, m)
	return (m + x) % n
}
