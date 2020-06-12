package leetcode



// 动态规划
func generate_yh(numRows int) [][]int {
	yh := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			yh[i] = append(yh[i], helper_yh(i, j, yh))
		}
	}
	return yh
}

func helper_yh(i, j int, yh [][]int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if j == 0 || i == j {
		return 1
	}
	return yh[i-1][j-1] + yh[i-1][j]
}
