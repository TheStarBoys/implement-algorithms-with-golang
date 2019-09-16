package leetcode

// 递归
func postorderTraversal145_0(root *TreeNode) []int {
	ans := []int{}
	return helper145_0(root, ans)
}

func helper145_0(root *TreeNode, ans []int) []int {
	if root == nil { return ans }
	ans = helper145_0(root.Left, ans)
	ans = helper145_0(root.Right, ans)
	ans = append(ans, root.Val)
	return ans
}
// dfs迭代
func postorderTraversal145_1(root *TreeNode) []int {
	ans := []int{}
	return dfs145_1(root, ans)
}
func dfs145_1(root *TreeNode, ans []int) []int {
	stack := make([]*TreeNode, 0)
	cur := root
	visited := make(map[*TreeNode]bool)
	for cur != nil || len(stack) != 0 {
		for !visited[cur] && cur != nil {
			stack = append(stack, cur)
			visited[cur] = true
			cur = cur.Left
		}
		if len(stack) == 0 {
			break
		}
		cur = stack[len(stack) - 1]
		if !visited[cur.Right] && cur.Right != nil {
			cur = cur.Right
		} else {
			ans = append(ans, cur.Val)
			stack = stack[:len(stack) - 1]
		}
	}
	return ans
}

// LeetCode官方题解
func postorderTraversal145_2(root *TreeNode) []int {
	ans := []int{}
	stack := make([]*TreeNode, 0)
	if root == nil { return ans }
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ans = append([]int{node.Val}, ans...)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return ans
}