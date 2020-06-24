package leetcode

/*
// dfs
func maxAreaOfIsland(grid [][]int) int {
	visited := map[string]bool{}
	var ans int
	area = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j, visited, grid)
				if ans < area {
					ans = area
				}
				area = 0
			}
		}
	}

	return ans
}

var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
	area int
)
func dfs(i, j int, visited map[string]bool, grid [][]int) {
	visited[helper(i, j)] = true
	grid[i][j] = 0
	area++
	for k := 0; k < 4; k++ {
		tmp_i := i + dx[k]
		tmp_j := j + dy[k]
		if !visited[helper(tmp_i, tmp_j)] && tmp_i >= 0 && tmp_i < len(grid) &&
			tmp_j >= 0 && tmp_j < len(grid[0]) && grid[tmp_i][tmp_j] == 1 {
			dfs(tmp_i, tmp_j, visited, grid)
		}
	}
}

func helper(i, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}

 */
