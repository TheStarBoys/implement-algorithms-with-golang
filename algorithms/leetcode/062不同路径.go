package leetcode

// 不同路径, 动态规划
func uniquePaths062_0(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	}
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		t := make([]int, m)
		t[m-1] = 1 // 这个位置的元素只能往下走这一种情况
		nums[n-i-1] = t
		for j := m-2; j >= 0; j-- {
			dofunc062_0(n-i-1, j, nums)
		}
	}
	return nums[0][0]
}

func dofunc062_0(n, m int, nums [][]int) {
	i, j := 0, nums[n][m+1] // i, j分别保存该节点右侧跟下侧的值
	if n < len(nums)-1 { // 如果不是最后一行
		i = nums[n+1][m]
	}
	nums[n][m] = i+j
}

// 不同路径Ⅱ, todo 题号位置，暂且放这里
// 时间复杂度太高了
func uniquePathsWithObstacles062(obstacleGrid [][]int) int {
	//count = 0 // init
	dfs062(obstacleGrid, 0, 0)
	return count062
}
// 机器人只有两个方向，向下或向右
//var dx = []int{0,1}
//var dy = []int{1,0}
var count062 int
func dfs062(grid [][]int, i, j int) {
	if i == len(grid) - 1 && j == len(grid[0]) - 1 {
		count062++
		return
	}
	for z := range dx {
		temp_i := i + dx[z]
		temp_j := j + dy[z]
		if 0 <= temp_i && temp_i < len(grid) && 0 <= temp_j && temp_j < len(grid[0]) &&
			grid[temp_i][temp_j] != 1 {
			dfs062(grid, temp_i, temp_j)
		}
	}
}