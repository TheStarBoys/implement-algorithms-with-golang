package leetcode

//
func countPrimes204_0(n int) int {
	table := make([]int, n)
	count := 0
	for i := 2; i < n; i++ {
		if table[i] == 0 {
			count++
			for j := i + i; j < n; j += i {
				table[j] = 1
			}
		}

	}
	return count
}
