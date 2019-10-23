package leetcode


// 944
func minDeletionSize944_0(A []string) int {
	D := []int{}
	for i := range A[0] { // 外层遍历列
		for j := range A[:len(A)-1] { // 内层遍历行
			if A[j][i] > A[j+1][i] { // 降序排列
				D = append(D, i)
				break
			}
		}
	}
	return len(D)
}

