package leetcode

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

// DFS
func inorderTraversal094_0(root *TreeNode) []int {
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
// 递归
func inorderTraversal094_1(root *TreeNode) []int {
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