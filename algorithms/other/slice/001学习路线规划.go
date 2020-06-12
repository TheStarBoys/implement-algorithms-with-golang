package slice

import "sort"

// 通过要学习的每门课，以及它的前置课程，来规划学习路线

//func main() {
//	for i, course := range topoSort(prereqs) {
//		fmt.Printf("%d:\t%s\n", i+1, course)
//	}
//}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// prereqs记录了每个课程的前置课程, 课程 --> 前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
