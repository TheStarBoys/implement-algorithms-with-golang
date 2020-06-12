package recursiveAlgorithm

// 1 + 2 + 3 +....+ (n - 1) + n
func sum(n int) int { // 没有处理错误情况，n必须是正整数
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}
