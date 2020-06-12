package leetcode

// 时间O(N) 空间O(N)
func rotateRight061_0(head *ListNode, k int) *ListNode {
	table := make([]*ListNode, 0)
	if head == nil { return head }
	for head != nil {
		table = append(table, head)
		head = head.Next
	}
	k = k % len(table)
	if k == 0 { return table[0] }
	table[len(table) - 1].Next = table[0] // 最后一个元素链接到头
	table[len(table) - k - 1].Next = nil // 第k个元素成为尾
	return table[len(table) - k]
}
// 时间O(N) 空间O(1)
func rotateRight061_10(head *ListNode, k int) *ListNode {
	// base cases
	if head == nil { return nil }
	if head.Next == nil { return head }

	// close the linked list into the ring
	old_tail := head
	n := 1
	for ; old_tail.Next != nil; n++ {
		old_tail = old_tail.Next
	}
	old_tail.Next = head

	// find new tail : (n - k % n - 1)th node
	// and new head : (n - k % n)th node
	new_tail := head
	for i := 0; i < n - k % n - 1; i++ {
		new_tail = new_tail.Next
	}
	new_head := new_tail.Next

	// break the ring
	new_tail.Next = nil
	return new_head
}