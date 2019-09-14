package leetcode

// 删除链表的倒数第N个节点019
func removeNthFromEnd019_0(head *ListNode, n int) *ListNode {
	if head == nil { return nil }
	table := make(map[int]*ListNode) // index --> *ListNode
	i := 0
	for ; head != nil; i++ {
		table[i] = head
		head = head.Next
	} // i == length
	if n > i { // 没有倒数N的节点
		return nil
	} else if n == i { // 删去头节点
		return table[0].Next
	}
	del := table[i - n] // waiting delete
	pre := table[i - n - 1]
	pre.Next = del.Next
	return table[0]
}

