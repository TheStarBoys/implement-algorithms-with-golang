package leetcode


/*
n = 1, out = 1 	1
n = 2, out = 2	1+1, 2
n = 3, out = 3	1+1+1, 1+2, 2+1
n = 4, out = 5	1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2

f(n) = f(n-1) + f(n-2)	n >= 3
 */

// recursiveAlgorithm
func climbStairs0(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs0(n-1) + climbStairs0(n-2)
}
// 带备忘录的递归
func climbStairs1(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	return helper(dp, n)
}

func helper(dp []int, n int) int{
	if n > 0 && dp[n] == 0 {
		dp[n] = helper(dp, n - 1) + helper(dp, n - 2)
	}
	return dp[n]
}

// 动态规划
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	sum := 0
	pre, cur := 1, 2
	for i := 2; i < n; i++ {
		sum = pre + cur
		pre, cur = cur, sum
	}
	return sum
}