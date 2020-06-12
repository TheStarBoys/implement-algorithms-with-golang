package leetcode


func removeElements203_0(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head =  head.Next
	}
	if head == nil { return nil }
	pre := head
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return head
}

