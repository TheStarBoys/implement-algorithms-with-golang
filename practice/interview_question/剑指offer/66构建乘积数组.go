package offer

// 暴力：超时
func constructArr_1(a []int) []int {
	b := make([]int, len(a))
	for i := range b {
		b[i] = 1
		for j := range a {
			if i == j { continue }
			b[i] *= a[j]
		}
	}
	return b
}

// 两个数组保存结果
func constructArr_2(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	// B[i] = A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]
	// B[i-1] = A[0]*A[1]*...*A[i-2]*A[i]*..*A[n-1]

	// dp1[i] -> A[0]*..*A[i] => dp1[i] = dp1[i-1] * A[i]
	// dp2[i] -> A[i]*..*A[n-1] => dp2[i] = dp2[i+1] * A[i]
	// b[i] = dp1[i-1] * dp2[i+1]
	n := len(a)
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp1[0] = a[0]
	dp2[n-1] = a[n-1]
	for i, j := 1, n-2; i < n && j >= 0; i, j = i+1, j-1 {
		dp1[i] = dp1[i-1] * a[i]
		dp2[j] = dp2[j+1] * a[j]
	}
	b := make([]int, n)
	b[0] = dp2[1]
	b[n-1] = dp1[n-2]
	for i := 1; i < n-1; i++ {
		b[i] = dp1[i-1] * dp2[i+1]
	}
	return b
}