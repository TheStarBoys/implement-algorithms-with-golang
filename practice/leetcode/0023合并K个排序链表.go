package leetcode

import "math"

// 类似归并，分治思想，大问题化小问题，小问题直接返回结果。
func mergeKLists0023_0(lists []*ListNode) *ListNode {
	n := len(lists)
	// 分治得到基本情况，直接解出答案
	if n == 0 { return nil }
	if n == 1 { return lists[0] }
	mid := n >> 1
	// 得到两个有序链表
	left := mergeKLists0023_0(lists[:mid])
	right := mergeKLists0023_0(lists[mid:])
	// 直接合并
	return merge0023_0(left, right)
}

func merge0023_0(l1, l2 *ListNode) *ListNode {
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
