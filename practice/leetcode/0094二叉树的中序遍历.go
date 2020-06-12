package leetcode

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

// 直接套用模版的dfs，不是太优雅
func inorderTraversal094_0(root *TreeNode) []int {
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
			stack = append(stack, node)
			visited[node] = true
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil && !visited[node.Right] {
			visited[node.Right] = true
			stack = append(stack, node.Right)
		}
	}
	return ans
}

// DFS
func inorderTraversal094_1(root *TreeNode) []int {
	res := []int{}
	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}
// recursiveAlgorithm
func inorderTraversal094_2(root *TreeNode) []int {
	ans := []int{}
	return helper094_1(root, ans)
}

func helper094_1(root *TreeNode, ans []int) []int {
	if root == nil { return ans }
	ans = helper094_1(root.Left, ans)
	ans = append(ans, root.Val)
	ans = helper094_1(root.Right, ans)
	return ans
}