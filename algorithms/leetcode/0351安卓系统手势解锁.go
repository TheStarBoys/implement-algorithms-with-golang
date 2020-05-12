package leetcode

// LeetCode 官方题解，回溯算法
func numberOfPatterns_0351(m int, n int) int {
	var res int
	for len := m; len <= n; len++ {
		res += calcPatterns_0351(-1, len)
		used = [9]bool{}
	}

	return res
}

var used = [9]bool{}
func calcPatterns_0351(last, len int) int {
	if len == 0 {
		return 1
	}
	sum := 0
	for i := 0; i < 9; i++ {
		if isValid_0351(i, last) {
			used[i] = true
			sum += calcPatterns_0351(i, len-1)
			used[i] = false
		}
	}

	return sum
}

func isValid_0351(indx, last int) bool {
	if used[indx] { return false }
	// 第一次使用
	if last == -1 { return true }
	// 国际象棋马的移动 或者 (在同一行 或 在同一列 的相邻元素)
	if (indx + last) % 2 == 1 {
		return true
	}
	// 对角线上两端的单元格，判断中间经过的单元格是否使用过
	mid := (indx + last)/2
	if mid == 4 { return used[mid] }
	// 对角线上的相邻单元格
	// indx%3 != last%3 用来确定他们不在同一列
	// indx/3 != last/3 用来确定他们不在同一行
	if (indx%3 != last%3) && (indx/3 != last/3) {
		return true
	}
	// 所有不相邻的单元格
	return used[mid]
}
