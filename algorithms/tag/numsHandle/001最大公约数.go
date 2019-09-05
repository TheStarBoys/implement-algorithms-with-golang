package numsHandle

// 欧几里得解法
func Euclid001(m, n int) int {
	// 使用欧几里得算法计算gcd(m, n)
	// 输入两个不全为0的非负整数
	// 输出他们的最大公约数
	if m == 0 && n == 0 {
		return -1
	}
	for n != 0 {
		m, n = n, m % n
	}
	return m
}
// 筛选法，找出<=n的所有质数
func Sieve(n int) []int {
	// 输入：一个正整数 n > 1
	if n <= 1 {
		return []int{}
	}
	A := make([]int, n + 1)
	for i := 2; i <= n; i++ {
		A[i] = 1
	}
	for p := 2; p * p <= n; p++ {
		if A[p] != 0 { // p没有被前面的步骤消去
			j := p * p
			for j <= n {
				A[j] = 0 // 标记为消去
				j += p
			}
		}
	}
	// 将A中剩余元素复制到质数数组中
	L := []int{}
	for i := range A {
		if A[i] != 0 {
			L = append(L, i)
		}
	}
	return L
}