package leetcode

// 相交链表160
func getIntersectionNode160_0(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { return nil }
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

