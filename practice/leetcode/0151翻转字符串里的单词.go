package leetcode

import "strings"

func reverseWords(s string) string {
	strs := []string{}
	prevIndx := 0
	// 过滤多余空格
	for {
		for ; prevIndx < len(s) && s[prevIndx] == ' '; prevIndx++ {}
		if prevIndx == len(s) {
			break
		}
		i := prevIndx
		for ; i < len(s) && s[i] != ' '; i++ {}
		strs = append(strs, s[prevIndx:i])
		prevIndx = i
	}
	// 反转
	reverseStrs(strs)

	return strings.Join(strs, " ")
}