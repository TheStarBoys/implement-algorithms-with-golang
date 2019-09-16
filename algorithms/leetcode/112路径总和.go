package leetcode

// ps: 这道题需要注意的坑在于 路径是从根节点到叶子节点
// 如果出现中途的和等于目标和不作数

// dfs隐式调用系统栈
func hasPathSum112_0(root *TreeNode, sum int) bool {
	return dfs112_0(root, 0, sum)
}
func dfs112_0(root *TreeNode, curSum, sum int) bool {
	if root == nil { return false }
	curSum += root.Val
	if root.Left == nil && root.Right == nil && curSum == sum { return true }
	if dfs112_0(root.Left, curSum, sum) { return true }
	return dfs112_0(root.Right, curSum, sum)
}

// 第一种代码可以简化成如下
func hasPathSum112_1(root *TreeNode, sum int) bool {
	if root == nil { return false }
	sum -= root.Val
	if root.Left == nil && root.Right == nil { return sum == 0 }
	return hasPathSum112_1(root.Left,sum) || hasPathSum112_1(root.Right, sum)
}

// dfs显示调用系统栈
func hasPathSum112_2(root *TreeNode, sum int) bool {
	if root == nil { return false }
	node_stack := make([]*TreeNode, 0)
	sum_stack := make([]int, 0)
	node_stack = append(node_stack, root)
	sum_stack = append(sum_stack, sum - root.Val)
	node := &TreeNode{}
	curr_sum := 0
	for len(node_stack) != 0 {
		node = node_stack[len(node_stack) - 1]
		node_stack = node_stack[:len(node_stack) - 1]

		curr_sum = sum_stack[len(sum_stack) - 1]
		sum_stack = sum_stack[:len(sum_stack) - 1]

		if node.Right == nil && node.Left == nil && curr_sum == 0 {
			return true
		}
		if node.Right != nil {
			node_stack = append(node_stack, node.Right)
			sum_stack = append(sum_stack, curr_sum - node.Right.Val)
		}
		if node.Left != nil {
			node_stack = append(node_stack, node.Left)
			sum_stack = append(sum_stack, curr_sum - node.Left.Val)
		}
	}
	return false
}






