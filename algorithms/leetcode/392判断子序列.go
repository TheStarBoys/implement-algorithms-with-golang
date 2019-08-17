package leetcode

import "strings"



func isSubsequence1(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i >= len(s)
}

func isSubsequence2(s string, t string) bool {
	for _, i := range s {
		tmp := strings.IndexRune(t, i)
		if tmp > -1 {
			t = t[tmp+1:]
		} else {
			return false
		}
	}
	return true
}
