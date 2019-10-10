package leetcode

import "sort"


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
