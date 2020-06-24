package offer

func maxDepth(root *TreeNode) int {
	if root == nil { return 0 }
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
