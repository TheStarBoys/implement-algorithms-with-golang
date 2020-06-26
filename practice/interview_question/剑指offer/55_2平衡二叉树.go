package offer

func isBalanced(root *TreeNode) bool {
	if root == nil { return true }
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil { return 0 }
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
	left_depth := maxDepth(root.Left)
	right_depth := maxDepth(root.Right)
	if left_depth >= right_depth && left_depth - right_depth > 1 ||
		right_depth - left_depth > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
