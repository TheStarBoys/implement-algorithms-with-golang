package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"aa","ab"}))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
	l, h := 0, minLen - 1
	for l <= h {
		mid := l + (h - l ) / 2
		if isCommonPrefix(l, mid, strs) {
			l = mid + 1
		}else {
			h = mid - 1
		}
	}
	return strs[0][:h + 1]
}
func isCommonPrefix(l, h int, strs []string) bool {
	prefix := strs[0][l : h + 1]
	for i := 1; i < len(strs); i++ {
		if strings.Index(strs[i][l : h + 1], prefix) == -1 {
			return false
		}
	}
	return true
}
