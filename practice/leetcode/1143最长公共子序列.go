package leetcode

// 回溯，会超时。包级变量有点多，就屏蔽了

/*

func longestCommonSubsequence(text1 string, text2 string) int {
	maxDist = math.MinInt64
	t1, t2 = text1, text2
	n, m = len(text1), len(text2)
	backtrack(0, 0, 0)
	return maxDist
}

var(
	maxDist int
	t1, t2 string
	n, m int
)

func backtrack(i, j, dist int) {
	if i == n || j == m {
		if dist > maxDist {
			maxDist = dist
		}
		return
	}
	if t1[i] == t2[j] {
		backtrack(i+1, j+1, dist+1)
	} else {
		backtrack(i+1, j, dist)
		backtrack(i, j+1, dist)
	}
}
 */

// 回溯，剪枝

/*
func longestCommonSubsequence(text1 string, text2 string) int {
    maxDist = math.MinInt64
    t1, t2 = text1, text2
    n, m = len(text1), len(text2)
    mem = make([][]int, n)
    for i := range mem {
        mem[i] = append(mem[i], make([]int, m)...)
    }
    backtrack(0, 0, 0)
    return maxDist
}

var(
    maxDist int
    t1, t2 string
    n, m int
    mem [][]int
)

func backtrack(i, j, dist int) {
    if i == n || j == m {
        if dist > maxDist {
            maxDist = dist
        }
        return
    }
    if mem[i][j] > 0 {
        if mem[i][j] >= dist {
            return
        }
    }
    mem[i][j] = dist
    if t1[i] == t2[j] {
        backtrack(i+1, j+1, dist+1)
    } else {
        backtrack(i+1, j, dist)
        backtrack(i, j+1, dist)
    }
}
 */

// 动态规划
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	if n * m == 0 {
		return 0
	}
	// 如果字母不同：dp(i, j) = max(dp(i-1, j), dp(i, j-1))
	// 如果字母相同：dp(i, j) = dp[i-1][j-1] + 1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = append(dp[i], make([]int, m)...)
	}
	for i := 0; i < n; i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else if i != 0 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 0; j < m; j++ {
		if text1[0] == text2[j] {
			dp[0][j] = 1
		} else if j != 0 {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n-1][m-1]
}