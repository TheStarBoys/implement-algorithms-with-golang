package leetcode

// 暴力 O(n^2)
func maxScoreSightseeingPair1014_0(A []int) int {
	var res int
	for j := 1; j < len(A); j++ {
		for i := 0; i < j; i++ {
			res = max(res, A[i] + A[j] + i - j)
		}
	}

	return res
}

//
func maxScoreSightseeingPair1014_1(A []int) int {
	var (
		ans int
		mx = A[0] + 0
	)
	for j := 1; j < len(A); j++ {
		ans = max(ans, mx + A[j] - j)
		mx = max(mx, A[j] + j)
	}

	return ans
}
