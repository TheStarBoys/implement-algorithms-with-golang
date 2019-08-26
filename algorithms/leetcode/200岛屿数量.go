package leetcode

// BFS
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}
var row, col int
func numIslands(grid [][]byte) int {
	row = len(grid)
	if row == 0{
		return 0
	}

	col = len(grid[0])
	count := 0
	for i:=0; i<row; i++{
		for j:=0; j<col; j++{
			if grid[i][j] == '1'{
				BFS(grid, i, j)
				count++
			}
		}
	}
	return count
}
// 广度优先搜索，查找出当前坐标，与之在垂直和水平方向相连的陆地
func BFS(grid [][]byte, i, j int){
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
