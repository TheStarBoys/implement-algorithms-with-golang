package leetcode

/*
// 回溯，超时
func numWays(n int, k int) int {
	N = n
	K = k
	count = 0
	for i := 0; i < k; i++ {
		backtrack(0, i, 1)
	}
	return count
}

var (
	N int
	K int
	count int
)

func backtrack(i, clr, preClrCnt int) {
	if i == N - 1 {
		count++
		return
	}
	for j := 0; j < K; j++ {
		if j == clr {
			if preClrCnt < 2 {
				backtrack(i+1, j, preClrCnt+1)
			}
		} else {
			backtrack(i+1, j, 1)
		}
	}
}
 */

func numWays0276_0(n int, k int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return k
	} else if n == 2 {
		return k * k
	}
	// 已知 i - 1 的涂色方案，求 i 的涂色方案
	// 若 i - 1 跟 i - 2 颜色相同，那么 i 有 k - 1 种选择
	// 若 i - 1 跟 i - 2 颜色不同，那么 i 有 k 种选择
	// 相同次数为 i - 1 跟 i - 2 颜色相同 情况的次数

	// n = 3, k = 3
	// | i | 相同次数 | 有效涂色方案
	// | 1 |    0    |      k = 3
	// | 2 |   k=3   |      k^2 = 9
	// | 3 |  9-3=6  |    (9-3)k + 3(k-1) = 24

	// 有效涂色方案，相同次数
	res, cnt := k*k, k
	for i := 3; i <= n; i++ {
		tmp_cnt := res - cnt
		res = tmp_cnt * k + cnt*(k-1)
		cnt = tmp_cnt
	}
	return res
}