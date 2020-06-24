package offer

import "math"

func printNumbers(n int) []int {
	ans := make([]int, int(math.Pow(10, float64(n)))-1)
	for i := 0; i < len(ans); i++ {
		ans[i] = i+1
	}
	return ans
}
