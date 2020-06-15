package leetcode

import "math"

// 回溯
func minDistance0072_0(word1 string, word2 string) int {
	if len(word1) == 0 && len(word2) == 0 {
		return 0
	}
	minDist0072_0 = math.MaxInt64
	backtrack0072_0(0, 0, 0, word1, word2)
	return minDist0072_0
}

var minDist0072_0 int
func backtrack0072_0(i, j, dist int, w1, w2 string) {
	if i == len(w1) || j == len(w2) {
		if i < len(w1) {
			dist += len(w1) - i
		}
		if j < len(w2) {
			dist += len(w2) - j
		}
		if dist < minDist0072_0 {
			minDist0072_0 = dist
		}
		return
	}
	if w1[i] == w2[j] { // 两个字符匹配
		backtrack0072_0(i+1, j+1, dist, w1, w2)
	} else {
		dist++
		backtrack0072_0(i+1, j, dist, w1, w2) // 删除a[i]或者b[j]前添加一个字符
		backtrack0072_0(i, j+1, dist, w1, w2) // 删除b[j]或者a[i]前添加一个字符
		backtrack0072_0(i+1, j+1, dist, w1, w2) // 将a[i]和b[j]替换为相同字符
	}
}

// 动态规划
func minDistance0072_1(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	// 有一个是空串
	if n * m == 0 {
		return n + m
	}

	// 初始化 dp
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = append(dp[i], make([]int, m+1)...)
	}
	for i := 0; i < n+1; i++ { // 初始化 第 0 列
		// 相当于对 word1 执行 i 次删除操作
		dp[i][0] = i
	}
	for j := 0; j < m+1; j++ { // 初始化 第 0 行
		// 相当于对 word1执行 j 次插入操作
		dp[0][j] = j
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			left := dp[i][j-1] + 1
			down := dp[i-1][j] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_down++
			}
			dp[i][j] = min(left, min(left_down, down))
		}
	}
	return dp[n][m]
}