package offer

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil { return true }
	if left == nil || right == nil { return false }
	if left.Val != right.Val { return false }
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
