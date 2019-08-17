package main

import (
	"fmt"
	"strings"
)

//type elem []int
func main() {
	str := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(str))
}

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
