package offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil { return nil }
	length := 0

	for curr := head; curr != nil; curr = curr.Next {
		length++
	}
	k = length - k
	for i := 0; i < k; i++ {
		head = head.Next
	}
	return head
}