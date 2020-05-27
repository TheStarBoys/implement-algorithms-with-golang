package linkedList

//func newPalindrome(s string) *LinkedList {
//	res := new(LinkedList)
//	cur := res
//	for i := range s {
//		cur.Data = s[i]
//		if i < len(s) - 1 {
//			cur.Next = new(LinkedList)
//			cur = cur.Next
//		}
//	}
//
//	return res
//}
//
//func isPalindrome(head *LinkedList) bool {
//	// 如果是回文串，那么正序跟反序遍历结果相同
//	newHead := new(LinkedList)
//
//	cur := head
//	// 把新链的生成跟反转同时进行
//	var pre *LinkedList
//	for cur != nil {
//		newHead.Data = cur.Data
//		if cur.Next != nil {
//			newHead.Next = new(LinkedList)
//		}
//		tmp := newHead.Next
//		newHead.Next = pre
//		pre = newHead
//		newHead = tmp
//		cur = cur.Next
//	}
//	newHead = pre
//	for newHead != nil && head != nil {
//		if newHead.Data.(byte) != head.Data.(byte) {
//			return false
//		}
//		newHead = newHead.Next
//		head = head.Next
//	}
//
//	return true
//}
