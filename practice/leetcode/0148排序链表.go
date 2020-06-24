package leetcode

import "math"

// 归并排序
// 时间复杂度：O(n log n), 递归需要 O(log n), 递归树节点内部合并需要 O(n)
func sortList0148_0(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { return head }
	// 快慢指针找到链表中点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow is median
	right := sortList0148_0(slow.Next)
	slow.Next = nil // 注意这里需要主动断掉后面的链
	left := sortList0148_0(head)
	return merge0148_0(left, right) // 合并
}

// 合并思路其实就是LeetCode 21 号题：合并两个有序链表
func merge0148_0(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	for l1 != nil || l2 != nil {
		left, right := math.MaxInt32, math.MaxInt32
		if l1 != nil {
			left = l1.Val
		}
		if l2 != nil {
			right = l2.Val
		}
		if left <= right {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	return head.Next
}
