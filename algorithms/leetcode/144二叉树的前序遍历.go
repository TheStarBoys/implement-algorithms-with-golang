package leetcode

// 递归解法，非常简单
func preorderTraversal144_0(root *TreeNode) []int {
	ans := []int{}
	ans = helper144_0(root, ans)
	return ans
}
func helper144_0(root *TreeNode, ans []int) []int {
	if root == nil { return ans }
	ans = append(ans, root.Val)
	ans = helper144_0(root.Left, ans)
	ans = helper144_0(root.Right, ans)
	return ans
}
// 深度优先搜索
func preorderTraversal144_1(root *TreeNode) []int {
	if root == nil { return []int{} }
	ans := []int{}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		cur := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ans = append(ans, cur.Val)
		if cur.Right != nil { // 注意这里。需要先遍历右子树，因为在出栈的时候，先进的会后出
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}

	}
	return ans
}
// 莫里斯遍历 如果是实时输出，空间复杂度为O(1)
// 而每个前驱节点会被访问两次，时间复杂度为O(N)
func preorderTraversal144_2(root *TreeNode) []int {
	ans := []int{}
	node := root
	for node != nil {
		if node.Left == nil {
			ans = append(ans, node.Val)
			node = node.Right
		} else {
			predecessor := node.Left
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				ans = append(ans, node.Val)
				predecessor.Right = node
				node = node.Left
			} else {
				predecessor.Right = nil
				node = node.Right
			}
		}
	}
	return ans
}













