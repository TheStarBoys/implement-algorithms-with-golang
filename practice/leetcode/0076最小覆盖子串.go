package leetcode

// 滑动窗口
func minWindow0076(s string, t string) string {
	// 输出有效性判断
	if len(s) < len(t) {
		return ""
	}
	hash := [256]int{}
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}

	// 如果包含T所有字符的最小子串刚好就是S的话，那么min的边界条件就是需要大于S的长度
	l, count, result, min := 0, len(t), "", len(s) + 1
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}
		// 如果是多余或不需要的元素，将移动左指针
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}
		if count == 0 && min > r - l + 1 {
			min = r - l + 1
			result = s[l:r+1]
		}
	}

	return result
}
