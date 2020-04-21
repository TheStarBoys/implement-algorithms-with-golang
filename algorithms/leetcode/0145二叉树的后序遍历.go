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
	if root == nil { return ans }
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	visited := make(map[*TreeNode]bool)
	visited[root] = true
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		node := cur.Left
		for node != nil && !visited[node] {
			visited[node] = true
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if (node.Left == nil || visited[node.Left]) &&
			(node.Right == nil || visited[node.Right]) {
			ans = append(ans, node.Val)
			stack = stack[:len(stack)-1] // 注意这里的坑，在输出值的时候，才出栈
		}
		if node.Right != nil && !visited[node.Right] {
			stack = append(stack, node.Right)
			visited[node.Right] = true
		}

	}
	return ans
}

// dfs 2
func postorderTraversal145_2(root *TreeNode) []int {
	ans := []int{}
	return dfs145_2(root, ans)
}
func dfs145_2(root *TreeNode, ans []int) []int {
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
func postorderTraversal145_3(root *TreeNode) []int {
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