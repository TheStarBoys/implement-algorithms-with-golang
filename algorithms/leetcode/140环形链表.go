package leetcode

func hasCycle140_0(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && slow != fast {
		if fast.Next == nil {return false}
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast == nil {return false}
	return true
}

