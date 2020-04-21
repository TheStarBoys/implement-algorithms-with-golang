package leetcode

import (
	"bytes"
	"sort"
)

// 排序
func isAnagram242_0(s string, t string) bool {
	b1, b2 := []byte(s), []byte(t)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	return bytes.Compare(b1, b2) == 0
}

// 哈希表
func isAnagram242_1(s string, t string) bool {
	if len(s) != len(t) { return false }
	table := make(map[byte]int)
	for i := range s {
		table[s[i]]++
	}
	for i := range t {
		if table[t[i]] == 0 {
			return false
		}
		table[t[i]]--
	}
	return true
}
