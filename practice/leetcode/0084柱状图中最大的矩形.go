package leetcode

// 暴力解
func largestRectangleArea(heights []int) int {
	// 以 i 结尾的区间，有一 j 满足 [j, i] 的面积最大
	// [j, i] 的面积为 minH * (i - j + 1), 其中minH 为区间内的最小高度
	// 如果知道 [j, i] 的minH，就可以得到[j - 1, i]的 minH，进而得到 其区间
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		minH := heights[i]
		for j := i; j >= 0; j-- {
			if heights[j] < minH {
				minH = heights[j]
			}
			area := minH * (i - j + 1)
			if maxArea < area {
				maxArea = area
			}
		}
	}

	return maxArea
}
