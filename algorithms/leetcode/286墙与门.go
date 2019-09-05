package leetcode

import "fmt"

// 算法时间复杂度过高，超时
func wallsAndGates286_0(rooms [][]int)  {
	for i := range rooms {
		for j := range rooms[0] {
			if rooms[i][j] != 1 << 31 - 1 {
				continue
			}
			rooms[i][j] = BFS286_0(rooms, i, j, 0)
		}
	}
}
//var dx = []int{0, 0, -1, 1}
//var dy = []int{-1, 1, 0, 0}
func BFS286_0(rooms [][]int, i, j, target int) int {
	queue := [][]int{}
	step := 0

	queue = append(queue, []int{i, j})

	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			curIndex := queue[0]
			if rooms[curIndex[0]][curIndex[1]] == 0 { // find a door
				return step
			}
			for j := 0; j < len(dx); j++ {
				x := curIndex[0] + dx[j]
				y := curIndex[1] + dy[j]
				if x < 0 || x >= len(rooms) || y < 0 || y >= len(rooms[0]) {
					continue
				}
				if rooms[x][y] == -1 {
					continue
				}
				queue = append(queue, []int{x, y})
			}
			queue = queue[1:]
		}
		step++
	}
	return 1 << 31 - 1
}

// 针对上面代码的优化，添加了visited等，依旧超时，通过61/62个测试用例
func wallsAndGates286_1(rooms [][]int)  {
	for i := range rooms {
		for j := range rooms[0] {
			if rooms[i][j] != 1 << 31 - 1 {
				continue
			}
			rooms[i][j] = BFS286_1(rooms, i, j, 0)
		}
	}
}
//var dx = []int{0, 0, -1, 1}
//var dy = []int{-1, 1, 0, 0}
func BFS286_1(rooms [][]int, i, j, target int) int {
	queue := [][]int{}
	visited := make(map[string]bool)
	visited[k([]int{i, j})] = true
	step := 0

	queue = append(queue, []int{i, j})

	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			curIndex := queue[0]
			for j := 0; j < len(dx); j++ {
				x := curIndex[0] + dx[j]
				y := curIndex[1] + dy[j]
				if x < 0 || x >= len(rooms) || y < 0 || y >= len(rooms[0]) {
					continue
				}
				if rooms[x][y] == -1 {
					continue
				}
				if rooms[x][y] == 0 {
					return step + 1
				}
				if !visited[k([]int{x, y})] {
					queue = append(queue, []int{x, y})
					visited[k([]int{x, y})] = true
				}
			}
			queue = queue[1:]
		}
		step++
	}
	return 1 << 31 - 1
}
func k(list []int) string {
	return fmt.Sprintf("%q", list)
}