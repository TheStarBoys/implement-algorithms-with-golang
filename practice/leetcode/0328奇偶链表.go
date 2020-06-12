package leetcode

// 自己想出来的原始解答，有不少的冗余
func oddEvenList328_0(head *ListNode) *ListNode {
	if head == nil { return nil }
	if head.Next == nil { return head }
	// 奇数头跟偶数头
	head1, head2 := head, head.Next
	// 当前奇数尾 跟 偶数尾
	curTail1, curTail2 := head, head.Next
	for curTail1 != nil && curTail2 != nil {
		if curTail2.Next == nil { return head1 }
		curTail1.Next = curTail2.Next // 奇数尾指向下一个奇数
		curTail1 = curTail1.Next // 下一个奇数成为新的奇数尾
		curTail2.Next = curTail2.Next.Next // 偶数尾指向下一个偶数
		curTail2 = curTail2.Next // 下一个偶数替换当前偶数尾
		curTail1.Next = head2 // 下一个奇数 指向偶数头
	}
	return head1
}

// 根据LeetCode官方题解进行的优化后的答案
func oddEvenList328_1(head *ListNode) *ListNode {
	if head == nil { return nil }
	if head.Next == nil { return head }
	// 偶数头
	head2 := head.Next
	// 当前奇数尾 跟 偶数尾
	curTail1, curTail2 := head, head.Next
	for curTail2 != nil && curTail2.Next != nil {
		curTail1.Next = curTail2.Next // 奇数尾指向下一个奇数
		curTail1 = curTail1.Next // 下一个奇数成为新的奇数尾
		curTail2.Next = curTail1.Next // 偶数尾指向下一个偶数
		curTail2 = curTail2.Next // 下一个偶数替换当前偶数尾
	}
	curTail1.Next = head2 // 最后一个奇数 指向偶数头
	return head
}