package leetcode

// 50 模拟
func myPow050_0(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1.0 / x
	}
	ans := float64(1)
	for i := 0; i < n; i++ {
		ans *= x
	}
	return ans
}

// 快速幂算法（递归）
func myPow050_1(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1.0 / x
	}
	return fastPow050_1(x, n)
}

func fastPow050_1(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	// A == x^(n/2)
	A := fastPow050_1(x, n/2)
	if n % 2 == 0 { // 偶数
		// A*A = x^n,
		return A*A
	} else { // 奇数
		// A*A*x = x^n, A == x^(n/2)
		return x*A*A
	}
}


