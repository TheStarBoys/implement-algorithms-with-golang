package leetcode

// 先转置矩阵，再反转每一行的元素
// 栗子：
//
// | 1, 2, 3 |		| 1, 4, 7 |		| 7, 4, 1 |
// | 4, 5, 6 |  =>	| 2, 5, 8 | =>	| 8, 5, 2 |
// | 7, 8, 9 |		| 3, 6, 9 |		| 9, 6, 3 |
func rotate048_0(matrix [][]int)  {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := range matrix {
		l, r := 0, len(matrix[0]) - 1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}
