package leetcode

func searchBST700_0(root *TreeNode, val int) *TreeNode {
	if root == nil { return nil }
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST700_0(root.Left, val)
	} else {
		return searchBST700_0(root.Right, val)
	}
}

