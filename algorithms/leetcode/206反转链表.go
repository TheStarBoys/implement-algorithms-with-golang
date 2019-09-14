package leetcode


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

