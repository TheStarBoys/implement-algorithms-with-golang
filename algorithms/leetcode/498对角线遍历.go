package leetcode
//思路来源注明：// todo 暂时没思考太明白，放一放
//作者：wat-2
//链接：https://leetcode-cn.com/problems/diagonal-traverse/solution/java-jie-ti-xin-si-lu-dai-ma-zui-shao-by-wat-2/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func findDiagonalOrder498_0(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	flag := 0
	ans := make([]int, m * n)
	for i := 0; i <= m + n; i++ {
		if i % 2 == 0 {
			for j := m - 1; j >= 0; j-- {
				if i - j >= 0 && i - j < n {
					ans[flag] = matrix[j][i - j]
					flag++
				}
			}
		} else {
			for j := 0; j < m; j++ {
				if i - j >= 0 && i - j < n {
					ans[flag] = matrix[j][i - j]
					flag++
				}
			}
		}
	}
	return ans
}
