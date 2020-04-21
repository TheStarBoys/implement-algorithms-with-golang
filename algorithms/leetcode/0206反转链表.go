package leetcode

// 迭代
func reverseList206_0(head *ListNode) *ListNode {
	var pre *ListNode // 如果用pre := &ListNode{}，默认pre是有值的，而赋值nil的话会报错
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// 递归
func reverseList206_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { return head }
	newHead := reverseList206_1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
