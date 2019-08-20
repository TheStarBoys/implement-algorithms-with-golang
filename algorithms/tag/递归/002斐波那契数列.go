package 递归

// 1, 1, 2, 3, 5, 8

// 求第n项的值
func fibo(n int) int{
	if n == 1 || n == 2 {
		return 1
	}
	return fibo(n - 2) + fibo(n - 1)
}
