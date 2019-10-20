package leetcode

func insertIntoBST701_0(root *TreeNode, val int) *TreeNode {
	if root == nil { return nil }
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{val, nil, nil}
			return root
		}
		root.Left = insertIntoBST701_0(root.Left, val)
	} else if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{val, nil, nil}
			return root
		}
		root.Right = insertIntoBST701_0(root.Right, val)
	}
	return root
}

// 代码简化
func insertIntoBST701_1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:val,
		}
	}


	if val < root.Val {
		root.Left = insertIntoBST701_1(root.Left, val)
	} else {
		root.Right = insertIntoBST701_1(root.Right, val)
	}
	return root
}