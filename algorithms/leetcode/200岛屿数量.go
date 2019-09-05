package leetcode

// BFS
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}
var row, col int
func numIslands200_0(grid [][]byte) int {
	row = len(grid)
	if row == 0{
		return 0
	}

	col = len(grid[0])
	count := 0
	for i:=0; i<row; i++{
		for j:=0; j<col; j++{
			if grid[i][j] == '1'{
				BFS200_0(grid, i, j)
				count++
			}
		}
	}
	return count
}
// 广度优先搜索，查找出当前坐标，与之在垂直和水平方向相连的陆地
func BFS200_0(grid [][]byte, i, j int){
	queue := make([]int, 0)
	queue = append(queue, i, j)
	grid[i][j] = '0'
	for len(queue) != 0{
		i, j := queue[0], queue[1]
		queue = queue[2:]
		for m:=0; m<4; m++{
			tmp_i := i + dx[m]
			tmp_j := j + dy[m]

			if 0<=tmp_i && tmp_i<row && 0<=tmp_j && tmp_j<col && grid[tmp_i][tmp_j] == '1'{
				grid[tmp_i][tmp_j] = '0'
				queue = append(queue, tmp_i, tmp_j)
			}
		}
	}
}

// DFS 深度优先搜索 解法
func numIslands200_1(grid [][]byte) int {
	count := 0
	visited := make(map[string]bool)
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				DFS200_1(grid, visited, i, j)
				count++
			}
		}
	}
	return count
}

//func k(list []int) string {
//	return fmt.Sprintf("%q", list)
//}

//var dx = []int{0, 0, -1, 1}
//var dy = []int{-1, 1, 0, 0}
func DFS200_1(grid [][]byte, visited map[string]bool, i, j int) {
	if grid[i][j] == '0' {
		return
	} else {
		grid[i][j] = '0'
	}
	for z := range dx {
		temp_i := dx[z] + i
		temp_j := dy[z] + j
		if temp_i >= 0 && temp_i < len(grid) &&
			0 <= temp_j && temp_j < len(grid[0]) &&
			!visited[k([]int{temp_i, temp_j})] {
			visited[k([]int{temp_i, temp_j})] = true
			DFS200_1(grid, visited, temp_i, temp_j)
		}
	}
}