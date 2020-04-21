package leetcode

import "sort"
// 贪心策略
func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	gi, si := 0, 0
	for si < len(s) && gi < len(g){
		if g[gi] <= s[si] {
			gi++
		}
		si++
	}
	return gi
}

func FindContentChildren(g []int, s []int) int {
	return findContentChildren(g, s)
}