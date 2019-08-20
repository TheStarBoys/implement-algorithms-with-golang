package 递归

// n! = n * (n-1) * (n-2)* ... * 1 (n > 0)
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return factorial(n - 1) * n
}
