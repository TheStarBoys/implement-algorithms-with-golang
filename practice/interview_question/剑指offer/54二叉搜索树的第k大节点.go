package offer

func kthLargest(root *TreeNode, k int) int {
	var inorder func(root *TreeNode)
	nums := []int{}
	inorder = func(root *TreeNode) {
		if root == nil { return }
		inorder(root.Left)
		nums = append(nums, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return nums[len(nums)-k]
}