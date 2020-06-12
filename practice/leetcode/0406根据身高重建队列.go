package leetcode

import "sort"
// todo 评分太低
func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 {
		return [][]int{}
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	rebuild := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		n := people[i][1]
		tmp := append([][]int{}, rebuild[n:]...)
		rebuild = append(rebuild[:n], people[i])
		rebuild = append(rebuild, tmp...)
	}
	return rebuild[:len(people)]
}
// golang 不支持这样封装一个函数来操作切片，外层切片会很诡异
func insert(rebuild [][]int, elem[]int, n int) {
	tmp := append([][]int{}, rebuild[n:]...)
	rebuild = append(rebuild[:n], elem)
	rebuild = append(rebuild, tmp...)
}
