package stringMatching

func KMPSearch(mainStr, patternStr string) int {
	if len(mainStr) < len(patternStr) {
		return -1
	}
	return kmp(mainStr, patternStr, len(mainStr), len(patternStr))
}

// n, m分别为主串和模式串长度
func kmp(mainStr, patternStr string, n, m int) int {
	next := getNexts(patternStr, m)
	j := 0
	for i := 0; i < n; i++ {
		// 当有好前缀 并且 此时主串和模式串出现坏字符时
		for j > 0 && mainStr[i] != patternStr[j] {
			// 从next数组中获取好前缀[0, j-1] 可匹配的最长前缀子串结尾下标
			j = next[j - 1] + 1
		}
		if mainStr[i] == patternStr[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}

	return -1
}
// next数组，前缀结束下标 : 最长可匹配前缀子串结尾字符下标
// 例如：模式串：ababacd
// 模式串前缀（好前缀候选） 前缀结尾字符下标 最长可匹配...下标
// a							0 		: 		-1
// ab							1 		: 		-1
// aba							2 		: 		 0
// abab							3 		: 		 1
// ababa						4 		: 		 2
// ababac						5 		: 		-1

// 以其中 abab 好前缀候选为例：
// 后缀	前缀		可匹配结尾下标
// b	a		-1 // 自己不能跟自己匹配
// ab	ab		 1
// bab	aba		-1
// abab	abab	-1 // 自己不能跟自己匹配
// 最长结尾下标 --> 1
func getNexts(patternStr string, m int) []int {
	next := make([]int, m)
	next[0] = -1
	k := -1
	// i 为好前缀候选结尾字符下标, next[m - 1] 的值其实肯定为 -1
	for i := 1; i < m; i++ {
		for k != -1 && patternStr[k + 1] != patternStr[i] {
			k = next[k]
		}
		// 如果相等，说明子串 [0, k+1] 就是 模式串[0, i]的最长可匹配前缀子串
		if patternStr[k + 1] == patternStr[i] {
			k++
		}
		next[i] = k
	}

	return next
}