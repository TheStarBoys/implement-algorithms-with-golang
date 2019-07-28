package main

import (
	"github.com/TheStarBoys/implement-data-structures-and-algorithms-with-golang/algorithms/leetcode"
)

func main() {
	l1 := &leetcode.ListNode{1,
		&leetcode.ListNode{2,
			&leetcode.ListNode{4, nil}}}
	l2 := &leetcode.ListNode{1,
		&leetcode.ListNode{3,
			&leetcode.ListNode{4, nil}}}
	leetcode.PrintNode(leetcode.MergeTwoLists(l1, l2))

}
