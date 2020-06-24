package leetcode

/*
// 回溯，超时
func minCost(costs [][]int) int {
    if len(costs) == 0 {
        return 0
    }
    min = math.MaxInt64
    n = len(costs)
    for i := 0; i < 3; i++ {
        backtrack(0, i, costs[0][i], costs)
    }
    return min
}

var (
    min int
    n int
)

func backtrack(i, j, cost int, costs [][]int) {
    if i == n - 1 {
        if min > cost {
            min = cost
        }
        return
    }
    for k := 0; k < 3; k++ {
        if k != j {
            backtrack(i+1, k, cost + costs[i+1][k], costs)
        }
    }
}
 */


/*
// 回溯，剪枝
func minCost(costs [][]int) int {
    if len(costs) == 0 {
        return 0
    }
    min = math.MaxInt64
    n = len(costs)
    mem = make([][]int, n)
    for i := 0; i < n; i++ {
        mem[i] = append(mem[i], make([]int, 3)...)
    }
    for i := 0; i < 3; i++ {
        backtrack(0, i, costs[0][i], costs)
    }
    return min
}

var (
    min int
    n int
    mem [][]int
)

func backtrack(i, j, cost int, costs [][]int) {
    if i == n - 1 {
        if min > cost {
            min = cost
        }
        return
    }
    if mem[i][j] > 0 {
        if mem[i][j] <= cost {
            return
        }
    }
    mem[i][j] = cost
    for k := 0; k < 3; k++ {
        if k != j {
            backtrack(i+1, k, cost + costs[i+1][k], costs)
        }
    }
}
 */

// 动态规划
func minCost0256_0(costs [][]int) int {
	for i := len(costs) - 2; i >= 0; i-- {
		costs[i][0] += min(costs[i+1][1], costs[i+1][2])
		costs[i][1] += min(costs[i+1][0], costs[i+1][2])
		costs[i][2] += min(costs[i+1][0], costs[i+1][1])
	}
	if len(costs) == 0 {
		return 0
	}

	return min(costs[0][0], min(costs[0][1], costs[0][2]))
}

// 动态规划，常数空间
func minCost0256_1(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	n := len(costs)
	r, g, b := costs[n-1][0], costs[n-1][1], costs[n-1][2]
	for i := n - 2; i >= 0; i-- {
		tmp_r, tmp_g, tmp_b := r, g, b
		r = costs[i][0] + min(tmp_g, tmp_b)
		g = costs[i][1] + min(tmp_r, tmp_b)
		b = costs[i][2] + min(tmp_r, tmp_g)
	}

	return min(r, min(g, b))
}