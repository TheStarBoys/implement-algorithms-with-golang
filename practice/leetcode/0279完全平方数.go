package leetcode

import "math"

func numSquares279_0(n int) int {
	absNums := []int{}
	for i := 1; i * i <= n; i++ {
		sqr := i * i
		if sqr == n {
			return 1
		}
		absNums = append(absNums, sqr)
	}

	// BFS
	queue := []int{}
	queue = append(queue, 0)
	step := 0

	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			if cur == n {
				return step
			}
			for _, v := range absNums {
				/*
				如果下面代码替换成这样, 时间复杂度过高，且会超出内存限制
				if num := cur + v; num <= n {
					queue = append(queue, num)
				}
				*/

				if num := cur + v; num == n {
					return step + 1
				} else if num < n {
					queue = append(queue, num)
				}
			}
			queue = queue[1:]
		}
		step++
	}
	return -1
}

// 比上一个方案多一个hashSet，效率提高非常多
func numSquares279_1(n int) int {
	absNums := []int{}
	for i := 1; i * i <= n; i++ {
		sqr := i * i
		if sqr == n {
			return 1
		}
		absNums = append(absNums, sqr)
	}

	// BFS
	queue := []int{}
	queue = append(queue, 0)
	used := make(map[int]bool)
	used[0] = true
	step := 0

	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			if cur == n {
				return step
			}
			for _, v := range absNums {
				if num := cur + v; num == n {
					return step + 1
				} else if !used[num] && num < n {
					queue = append(queue, num)
					used[num] = true
				}
			}
			queue = queue[1:]
		}
		step++
	}
	return -1
}

type Node297 struct {
	val int
	step int
}
// 这个方案，减少了获取从1-n所有的完全平方数的过程
func numSquares297_2(n int) int {
	queue := []Node297{}
	queue = append(queue, Node297{n, 0})
	visited := make([]bool, n + 1)

	for len(queue) != 0 {
		curNode := queue[0]
		num, step := curNode.val, curNode.step
		queue = queue[1:]

		for i := 1; ; i++ {
			a := num - i * i
			if a < 0 {
				break
			}
			if a == 0 {
				return step + 1
			}
			if !visited[a] {
				queue = append(queue, Node297{a, step + 1})
				visited[a] = true
			}
		}
	}
	return -1
}

// 这个方案，不需要额外的结构体
func numSquares297_3(n int) int {
	if n <= 0 { return 0 }
	if sqrt := int(math.Sqrt(float64(n))); sqrt * sqrt == n {
		return 1
	}
	count := 0
	queue := make([]int, 0)
	queue = append(queue, n)
	visited := make(map[int]bool)
	visited[n] = true

	for len(queue) > 0 {
		length := len(queue)
		count++
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			for j := 1; j*j <= cur; j++ {
				if tmp := j*j; !visited[cur - tmp] {
					if cur == tmp { return count }
					queue = append(queue, cur - tmp)
					visited[cur - tmp] = true
				}
			}
		}
	}
	return -1
}

// index: 0		1	2	3	4	..	n
// i*i  : 0		1	4	9	16		n*n
// 1	: 0		1
// 2	: 0		2
// 3	: 0 	3
// 4	: 0		0	1
// 5	: 0		1	1
// 6	: 0 	2	1
// 7	: 0 	3	1
// 8	: 0		0	2
// 9	: 0 	0	0	1
// 10 	: 0		1	0	1
// 11	: 0		2	0	1
// 12	: 0		3	0	1
// 13	: 0		0	1	1
// 14	: 0		1	1	1
// 15	: 0 	2	1	1
// 16	: 0 	0	0	0	1

// 贪心思想，每次拿能拿到的最大的数值。WARNING,这个方法是错误的
// 例如n=12的时候，输出为4(12=9+1+1+1)。而正确答案为3, 12= 4+4+4
func numSquares297_4(n int) int {
	if n <= 0 { return 0 }
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt * sqrt == n {
		return 1
	}
	table := make([]int, sqrt+1)
	for n > 0 {
		if n < 4 {
			table[1] = n
			break
		}
		sqrt := int(math.Sqrt(float64(n)))
		table[sqrt]++
		n = n - sqrt * sqrt // 15 - 3^2 = 6
	}
	ans := 0
	for _, v := range table {
		ans += v
	}
	return ans
}

