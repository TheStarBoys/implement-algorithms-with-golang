package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	tmp := head.Next
	head.Next = tmp.Next
	tmp.Next = head
	head = tmp
	head.Next.Next = swapPairs(head.Next.Next)
	return head
}
