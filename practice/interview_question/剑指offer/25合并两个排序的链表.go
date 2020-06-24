package offer

import "math"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	for l1 != nil || l2 != nil {
		left, right := math.MaxInt32, math.MaxInt32
		if l1 != nil {
			left = l1.Val
		}
		if l2 != nil {
			right = l2.Val
		}
		if left <= right {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	return head.Next
}
