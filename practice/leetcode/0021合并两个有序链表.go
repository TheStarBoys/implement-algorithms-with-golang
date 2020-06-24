package leetcode

import (
	"math"
)


// 第一种普通解法
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	head := &ListNode{}
//	tail := &ListNode{}
//	// 1. 对空链表做出处理
//	if l1 == nil {
//		return l2
//	} else if l2 == nil {
//		return l1
//	}
//	// 2. 找出第一个头结点
//	if l1.Val < l2.Val {
//		head = l1
//		l1 = l1.Next
//	} else {
//		head = l2
//		l2 = l2.Next
//	}
//	// 3. 不断更新tail
//	tail = head
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			tail.Next = l1
//			l1 = l1.Next
//		} else {
//			tail.Next = l2
//			l2 = l2.Next
//		}
//		tail = tail.Next
//	}
//	// 4. 把非空的那个链表链接到tail
//	if l1 != nil {
//		tail.Next = l1
//	} else if l2 != nil {
//		tail.Next = l2
//	}
//	return head
//}
// 递归思想
func mergeTwoLists0021_0(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		head = l1
		head.Next = mergeTwoLists0021_0(l1.Next, l2)
	}else {
		head = l2
		head.Next = mergeTwoLists0021_0(l1, l2.Next)
	}
	return head
}
// 解法2：循环遍历，O(n) n为两个链表中最小的长度
// 申请了额外的节点，空间复杂度为O(n)
func mergeTwoLists0021_1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	}
	head := &ListNode{}
	tmp := head
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			tmp.Val = l2.Val
			l2 = l2.Next
		}else {
			tmp.Val = l1.Val
			l1 = l1.Next
		}
		if l1 == nil || l2 == nil {
			break
		}
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	}else if l2 != nil {
		tmp.Next = l2
	}
	return head
}

// 利用链表本身的节点 O(1)空间复杂度
func mergeTwoLists0021_2(l1 *ListNode, l2 *ListNode) *ListNode {
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