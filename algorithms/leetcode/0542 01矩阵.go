package leetcode

// 简单bfs就做了，因为是最短路径，dfs不适合
func updateMatrix542_0(matrix [][]int) [][]int {
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] != 0 {
				bfs542_0(matrix, i, j)
			}
		}
	}
	return matrix
}
//var dx = []int{0, 0, -1, 1}
//var dy = []int{-1, 1, 0, 0}
func bfs542_0(matrix [][]int, i, j int) {
	visited := make(map[string]bool)
	visited[k([]int{i, j})] = true
	step := 0
	queue := make([][]int, 0)
	queue = append(queue, []int{i, j})

	for len(queue) != 0 {
		length := len(queue)
		for h := 0; h < length; h++ {
			cur := queue[0]
			queue = queue[1:]
			for z := range dx {
				temp_i := cur[0] + dx[z]
				temp_j := cur[1] + dy[z]
				if 0 <= temp_i && temp_i < len(matrix) && 0 <= temp_j && temp_j < len(matrix[0]) &&
					!visited[k([]int{temp_i, temp_j})] {
					if matrix[temp_i][temp_j] == 0 {
						matrix[i][j] = step + 1
						return
					} else {
						queue = append(queue, []int{temp_i, temp_j})
						visited[k([]int{temp_i, temp_j})] = true
					}
				}
			}
		}
		step++
	}
}
//func k(list []int) string {
//	return fmt.Sprintf("%q", list)
//}
