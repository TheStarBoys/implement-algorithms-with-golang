package leetcode

// 解法一：先得到链表长度，再递归
func reverseKGroup025_0(head *ListNode, k int) *ListNode {
	length := 0
	tmp := head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	return reverse025_0(head, length, k)
}
func reverse025_0(head *ListNode, length, k int) *ListNode {
	if length < k {
		return head
	}
	var pre *ListNode
	tail := head  // 1
	var tmp *ListNode
	for j := 0; j < k; j++ {
		tmp = head.Next // 2
		head.Next = pre //  1 --> nil
		pre = head
		head = tmp
	}
	// 2 -- > 1 -- > nil
	tail.Next = reverse025_0(tmp, length - k, k) // 2 -- > 1 -- > 4 -- > 3
	return pre
}