package offer

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil { return nil }
	if head.Val == val { return head.Next }
	var prev *ListNode
	curr := head
	for ; curr != nil && curr.Val != val; prev, curr = curr, curr.Next {}
	if curr == nil { return head }
	prev.Next = curr.Next
	return head
}
