package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 递归解法
func swapPairs024_0(head *ListNode) *ListNode {
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
	head.Next.Next = swapPairs024_0(head.Next.Next)
	return head
}
// 遍历解法
func swapPairs024_1(head *ListNode) *ListNode {
	newHead := head
	for head != nil && head.Next != nil {
		tmp := &ListNode{head.Val, head.Next.Next}
		head.Val = head.Next.Val
		head.Next = tmp
		head = head.Next.Next
	}
	return newHead
}