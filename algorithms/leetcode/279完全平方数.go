package leetcode

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
