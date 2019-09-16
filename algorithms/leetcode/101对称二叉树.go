package leetcode

// 递归
func isSymmetric101_0(root *TreeNode) bool {
	return isMirror101_0(root, root)
}
func isMirror101_0(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil { return true }
	if t1 == nil || t2 == nil { return false }
	return t1.Val == t2.Val &&
		isMirror101_0(t1.Right, t2.Left) &&
		isMirror101_0(t1.Left, t2.Right)
}
// 迭代
func isSymmetric101_1(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root, root}...)

	for len(queue) != 0 {
		t1 := queue[0]
		t2 := queue[1]
		queue = queue[2:]
		if t1 == nil && t2 == nil { continue }
		if t1 == nil || t2 == nil { return false }
		if t1.Val != t2.Val { return false }
		queue = append(queue, []*TreeNode{t1.Right, t2.Left}...)
		queue = append(queue, []*TreeNode{t2.Right, t1.Left}...)
	}
	return true
}