package leetcode

// 不同路径Ⅱ
// 时间复杂度太高了
func uniquePathsWithObstacles063_0(obstacleGrid [][]int) int {
	//count = 0 // init
	dfs063_0(obstacleGrid, 0, 0)
	return count063_0
}

// 机器人只有两个方向，向下或向右
//var dx = []int{0,1}
//var dy = []int{1,0}
var count063_0 int

func dfs063_0(grid [][]int, i, j int) {
	if i == len(grid)-1 && j == len(grid[0])-1 {
		count063_0++
		return
	}
	for z := range dx {
		temp_i := i + dx[z]
		temp_j := j + dy[z]
		if 0 <= temp_i && temp_i < len(grid) &&
			0 <= temp_j && temp_j < len(grid[0]) &&
			grid[temp_i][temp_j] != 1 {
			dfs063_0(grid, temp_i, temp_j)
		}
	}
}

// 动态规划
// 时间O(m*n) 空间O(1)
func uniquePathsWithObstacles063_1(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	r, c := len(obstacleGrid), len(obstacleGrid[0])
	obstacleGrid[0][0] = 1
	for i := 1; i < c; i++ {
		if obstacleGrid[0][i] == 1 {
			obstacleGrid[0][i] = 0
		} else {
			obstacleGrid[0][i] = obstacleGrid[0][i-1]
		}
	}
	for i := 1; i < r; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		}
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[r-1][c-1]
}