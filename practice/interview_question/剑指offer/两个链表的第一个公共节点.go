package offer

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { return nil }
	currA, currB := headA, headB
	for !(currA == nil && currB == nil) && currA != currB {
		if currA == nil {
			currA = headB
		} else {
			currA = currA.Next
		}
		if currB == nil {
			currB = headA
		} else {
			currB = currB.Next
		}
	}
	return currA
}
