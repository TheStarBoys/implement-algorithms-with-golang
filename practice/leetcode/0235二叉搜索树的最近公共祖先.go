package leetcode

func lowestCommonAncestor235_0(root, p, q *TreeNode) *TreeNode {
	left := 0
	right := 0
	mid := 0

	if p.Val < root.Val {
		left++
	} else if p.Val > root.Val {
		right++
	} else {
		mid++
	}

	if q.Val < root.Val {
		left++
	} else if q.Val > root.Val {
		right++
	} else {
		mid++
	}

	if mid > 0 || left == 1 && right == 1 { return root }
	if left > 1 {
		return lowestCommonAncestor235_0(root.Left, p, q)
	} else { // right > 1
		return lowestCommonAncestor235_0(root.Right, p, q)
	}
}

// 递归改成迭代
func lowestCommonAncestor235_1(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}
