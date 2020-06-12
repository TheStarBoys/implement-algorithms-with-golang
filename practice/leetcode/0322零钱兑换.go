package leetcode

import (
	"math"
	"sort"
)


// 贪心, is go die
func coinChange322_0(coins []int, amount int) int {
	sort.Ints(coins)
	p := len(coins) - 1
	count := 0
	for amount > 0 && p >= 0 {
		if tmp := amount - coins[p]; tmp >= 0 {
			count++
			amount = tmp
		} else {
			p--
		}
	}
	if amount == 0 { return count }
	return -1
}

// bfs
func coinChange322_1(coins []int, amount int) int {
	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	}
	// bfs
	sort.Ints(coins)
	queue := []int{} // curAmount
	queue = append(queue, amount)
	visited := make(map[int]bool)
	visited[amount] = true
	count := 0
	for len(queue) > 0 {
		length := len(queue)
		count++
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

			for p := len(coins) - 1; p >= 0; p-- {
				if tmp := cur - coins[p]; tmp == 0 {
					return count
				} else if tmp > 0 && !visited[tmp] {
					visited[tmp] = true
					queue = append(queue, tmp)
				}
			}
		}
	}
	return -1
}

// recursiveAlgorithm， 超时
func coinChange322_2(coins []int, amount int) int {
	// F(S) = F(S-C) + 1
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	min := math.MaxInt64
	for _, C := range coins {
		if amount - C < 0 { continue }
		subProb := coinChange322_2(coins, amount-C)
		if subProb == -1 { continue } // 子问题无解
		min = getMin322(min, subProb)
	}
	if min == math.MaxInt64 { return -1 }
	return min + 1
}

func getMin322(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

// 带备忘录的递归
func coinChange322_3(coins []int, amount int) int {
	// F(S) = F(S-C) + 1
	memo := make([]int, amount + 1)
	memo[0] = 0
	return dp322_3(coins, memo, amount)
}

func dp322_3(coins []int, memo []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount != 0 && memo[amount] == 0 {
		min := math.MaxInt64
		for _, C := range coins {
			if amount - C < 0 { continue }
			subProb := dp322_3(coins, memo, amount-C)
			if subProb == -1 { continue }
			min = getMin322(min, subProb)
		}
		if min == math.MaxInt64 {
			memo[amount] = -1
		} else {
			memo[amount] = min + 1
		}
	}
	return memo[amount]
}

// 递归改循环
func coinChange322_4(coins []int, amount int) int {
	// F(S) = F(S-C) + 1
	memo := make([]int, amount + 1)
	for i := range memo {
		memo[i] = amount + 1
	}
	memo[0] = 0
	for i := 1; i <= amount;i++ {
		for _, C := range coins {
			if i >= C {
				memo[i] = getMin322(memo[i-C]+1, memo[i])
			}
		}
	}
	if memo[amount] > amount {
		return -1
	}
	return memo[amount]
}

