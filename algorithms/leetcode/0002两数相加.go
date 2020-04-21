package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head
	for l1 != nil || l2 != nil {
		x, y, carry := 0, 0, 0
		if l1 != nil { x = l1.Val }
		if l2 != nil { y = l2.Val }
		val := x + y + tmp.Val
		if val >= 10 { carry = 1 }
		tmp.Val = val % 10
		if (l1 != nil && l1.Next != nil) || (l2 != nil && l2.Next != nil) || carry == 1 {
			tmp.Next = &ListNode{carry, nil}
			tmp = tmp.Next
		}
		if l1 != nil { l1 = l1.Next }
		if l2 != nil { l2 = l2.Next }
	}
	return head
}


