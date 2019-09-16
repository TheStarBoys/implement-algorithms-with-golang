package leetcode


func isPalindrome234_0(head *ListNode) bool {
	if head == nil || head.Next == nil { return true }
	slow, fast := head.Next, head.Next.Next // 快慢指针找到链表中点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var pre *ListNode
	var next *ListNode
	for head != slow { // 反转链表前一部分
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	if fast != nil { slow = slow.Next } // 如果是奇数个节点，则忽略后半部分的第一个节点
	for pre != nil { // 回文检验
		if pre.Val != slow.Val { return false }
		pre = pre.Next
		slow = slow.Next
	}
	return true
}
