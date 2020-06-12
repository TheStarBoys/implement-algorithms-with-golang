package leetcode

// 不同路径, 动态规划
func uniquePaths062_0(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	}
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		t := make([]int, m)
		t[m-1] = 1 // 这个位置的元素只能往下走这一种情况
		nums[n-i-1] = t
		for j := m-2; j >= 0; j-- {
			dofunc062_0(n-i-1, j, nums)
		}
	}
	return nums[0][0]
}

func dofunc062_0(n, m int, nums [][]int) {
	i, j := 0, nums[n][m+1] // i, j分别保存该节点右侧跟下侧的值
	if n < len(nums)-1 { // 如果不是最后一行
		i = nums[n+1][m]
	}
	nums[n][m] = i+j
}

// 动态规划 dp[i][j] = dp[i][j-1] + dp[i-1][j]
// 时间O(m * n) 空间O(m * n)
func uniquePaths062_1(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[n-1][m-1]
}
// 优化空间复杂度O(2n)
func uniquePaths062_2(m int, n int) int {
	pre := make([]int, n) // 上一行的所有元素
	cur := make([]int, n) // 当前行的所有元素
	for i := 0; i < n; i++ { pre[i] = 1 }
	for i := 0; i < n; i++ { cur[i] = 1 }
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = cur[j-1] + pre[j]
		}
		copy(pre, cur)
	}
	return pre[n-1]
}
// 优化空间复杂度O(n)
func uniquePaths062_3(m int, n int) int {
	cur := make([]int, n)
	for i := 0; i < n; i++ { cur[i] = 1 }
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// cur[j] 实际存储的是dp[i-1][j]
			// cur[j-1] 实际存储的是dp[i][j-1]
			cur[j] += cur[j-1]
		}
	}
	return cur[n-1]
}
