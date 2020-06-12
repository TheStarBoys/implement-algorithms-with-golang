package leetcode

// LeetCode官方题解golang版：模拟
func spiralOrder054_0(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 0 {
		return ans
	}
	R, C := len(matrix), len(matrix[0])
	var seen = make(map[string]bool)	// 存储是否被访问过
	// 通过dr跟dc来获取下一步走的方向
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}
	// r,c 表示访问的下标, di表示方向
	r, c, di := 0, 0, 0
	for i := 0; i < R * C; i++ {
		ans = append(ans, matrix[r][c])
		seen[k([]int{r, c})] = true
		// 获取下一步的下标
		cr := r + dr[di]
		cc := c + dc[di]
		if 0 <= cr && cr < R && 0 <= cc && cc < C &&
			!seen[k([]int{cr, cc})] { // 下标没越界，且没访问过
			r = cr
			c = cc
		} else { // 换个方向访问下一个节点
			di = (di + 1) % 4
			r += dr[di]
			c += dc[di]
		}
	}
	return ans
}

//func k(list []int) string {
//	return fmt.Sprintf("%q", list)
//}
