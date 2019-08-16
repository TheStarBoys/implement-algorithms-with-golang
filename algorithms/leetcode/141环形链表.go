package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	l1, l2 := head, head.Next
	for l1 != nil && l2 != nil && l2.Next != nil {
		if l1 == l2 {
			return true
		}
		l1 = l1.Next
		l2 = l2.Next.Next
	}
	return false
}
