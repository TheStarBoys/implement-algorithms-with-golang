package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	tmp.Val = l1.Val
	tmp.Next = l2
	return nil
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists(l1, l2)
}

func PrintNode(l *ListNode) {
	for l.Next == nil {
		fmt.Print(l.Val)
		l = l.Next
	}
}