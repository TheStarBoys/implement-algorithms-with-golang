package leetcode

import "strings"

func partitionLabels(S string) []int {
	ps := make([]int, 0)
	s := make(map[byte]int)
	begin := 0
	substr := ""
	for i := range S {
		if _, ok := s[S[i]]; !ok { // 如果map中没有该元素， 该步骤是为了保证去重，用空间换时间
			substr += string(S[i])
			s[S[i]] = 0
		}
		if !strings.ContainsAny(S[i+1:], substr) { // 不包含该元素
			ps = append(ps, i-begin+1)
			begin = i + 1
			substr = ""
			s = make(map[byte]int)
		}
	}
	return ps
}

// 自己实现的算法 效率较低
func partitionLabels0(S string) []int {
	slice := []int{}
	if len(S) == 0 {
		return slice
	}
	sub := string(S[0])
	for {
		for {
			i := contain(S, sub)
			if i == -1 {
				break
			}
			sub = string(S[:i+1])
		}
		slice = append(slice, len(sub))
		if len(sub) >= len(S) {
			return slice
		}
		S = S[len(sub):]
		sub = string(S[0])
	}
}
// 找到另外片段中相同字母出现的最后位置，不存在则返回负一
func contain(str, sub string) int{
	for i, _ := range sub {
		for j := len(str)-1; j >= len(sub); j-- { // 找到最后出现的位置
			if sub[i] == str[j] {
				return j
			}
		}
	}
	return -1
}
