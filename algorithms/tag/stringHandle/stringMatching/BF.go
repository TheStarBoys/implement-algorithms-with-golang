package stringMatching

// Brute Force
// 时间复杂度：O(n*m) n 为主串长度，m 为模式串长度
func BFSearch(str, sub string) int {
	n, m := len(str), len(sub)
	Flag:
	for i := 0; i < n - m + 1; i++ {
		// cur 为主串和模式串对齐的下标
		for cur, j := i, 0; j < m; cur, j = cur+1, j+1 {
			if str[cur] != sub[j] {
				continue Flag
			}
		}
		return i
	}

	return -1
}

//BF search pattern index, return the first match subs start index
func bfSearch(main string, pattern string) int {

	//defensive
	if len(main) == 0 || len(pattern) == 0 || len(main) < len(pattern) {
		return -1
	}

	for i := 0; i <= len(main)-len(pattern); i++ {
		subStr := main[i : i+len(pattern)]
		if subStr == pattern {
			return i
		}
	}

	return -1
}
