package offer

func reversePrint(head *ListNode) []int {
	if head == nil { return nil }
	return append(reversePrint(head.Next), head.Val)
}
