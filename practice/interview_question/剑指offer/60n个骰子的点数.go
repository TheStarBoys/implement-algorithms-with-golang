package offer

import "math"

// 回溯：超时
func twoSum_1(n int) []float64 {
	table := make(map[int]int)
	var backtrack func(sum int, n int)
	backtrack = func(sum, n int) {
		if n == 0 {
			table[sum]++
			return
		}
		for i := 1; i <= 6; i++ {
			backtrack(sum + i, n - 1)
		}
	}
	backtrack(0, n)
	total := math.Pow(6, float64(n))
	var ans []float64
	for i := n; i <= 6*n; i++ {
		ans = append(ans, float64(table[i]) / total)
	}
	return ans
}
