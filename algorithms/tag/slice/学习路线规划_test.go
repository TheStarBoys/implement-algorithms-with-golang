package slice

import (
	"testing"
	"fmt"
)

func TestTopoSort(t *testing.T) {
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
	courses := topoSort(prereqs)
	for i, course := range courses {
		if i == len(courses) - 1 {
			fmt.Printf("%s\n", course)
			break
		}
		fmt.Printf("%s --> ", course)
	}
}
