package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
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
func mergeTwoLists (l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		head = l1
		head.Next = mergeTwoLists(l1.Next, l2)
	}else {
		head = l2
		head.Next = mergeTwoLists(l1, l2.Next)
	}
	return head
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists(l1, l2)
}

func PrintNode(l ListNode) {
	for l.Next != nil {
		fmt.Print(l.Val)
		l = *l.Next
	} // 此时获得了最后一个l节点
	fmt.Print(l.Val)
	fmt.Println()
}