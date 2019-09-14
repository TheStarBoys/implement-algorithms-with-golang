package leetcode

// 环形链表贰 142
func detectCycle142_0(head *ListNode) *ListNode {
	// 如果s/f在入环的第一个节点， 那么再次相遇的位置，依旧是该节点
	if head == nil { return nil }
	s, f := head, head // 快慢指针
	tmp := head // 保存可能是入环节点的值
	for {
		if f.Next == nil || f.Next.Next == nil { return nil }
		s = s.Next
		f = f.Next.Next
		if s == f { // 说明有环
			if tmp == s { return tmp }
			// tmp不是入环节点
			s = tmp.Next
			f = s
			tmp = s
		}
	}
}

