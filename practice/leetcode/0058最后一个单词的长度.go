package leetcode


func lengthOfLastWord058_0(s string) int {
	if s == "" { return 0 }

	space := 0
	// 去掉末尾多余的空格，例如"a   "
	for i := len(s) - 1; i >= 0 && s[i] == ' '; i-- {
		space++
	}
	// 寻找单词之间的间隔，如果没找到，说明只有一个单词
	for i := len(s) - 1 - space; i >= 0; i-- {
		if s[i] == ' ' {
			// 11 - 5 -1 = 5
			return len(s) - 1 - space - i
		}
	}
	return len(s) - space
}

