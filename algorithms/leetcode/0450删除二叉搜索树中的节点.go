package leetcode

// 错误的解法，注意为什么会错
/*
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil { return nil }
	if key == root.Val {
		if root.Left == nil && root.Right == nil { return nil }
		// 错误的版本
		// 由于root.Left的可能还有孩子，所以这样直接连接会出问题
		if root.Left != nil {
			root.Left.Right = root.Right
			return root.Left
		}
		if root.Right != nil {
			root.Right.Left = root.Left
			return root.Right
		}
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
*/

// 先找到需要删除的节点，如果该节点没有孩子，直接删除
// 如果有孩子，则需要将一个孩子插入在另一个孩子节点里
func deleteNode450_0(root *TreeNode, key int) *TreeNode {
	if root == nil { return nil }
	if key == root.Val {
		// 如果是个叶子节点
		if root.Left == nil && root.Right == nil { return nil }
		// 如果左孩子不为空，把右孩子插入到左孩子里，并且满足二叉搜索树特性
		if root.Left != nil {
			return insertNode450_0(root.Left, root.Right)
		}
		// 如果右孩子不为空，把左孩子插入到右孩子里，并且满足二叉搜索树特性
		if root.Right != nil {
			return insertNode450_0(root.Right, root.Left)
		}
	} else if key < root.Val {
		root.Left = deleteNode450_0(root.Left, key)
	} else {
		root.Right = deleteNode450_0(root.Right, key)
	}
	return root
}

func insertNode450_0(root, insert *TreeNode) *TreeNode {
	if insert == nil { // 注意插入的节点可能为nil，因为删除函数没有做相应判断
		return root
	}
	if root == nil {
		return insert
	}
	if insert.Val < root.Val {
		root.Left = insertNode450_0(root.Left, insert)
	} else {
		root.Right = insertNode450_0(root.Right, insert)
	}
	return root
}