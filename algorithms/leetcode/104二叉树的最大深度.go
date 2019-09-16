package leetcode

import "math"

// 自顶向下的递归
func maxDepth104_0(root *TreeNode) int {
	helper104_0(root, 1)
	return answer
}
var answer int
func helper104_0(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		answer = int(math.Max(float64(answer), float64(depth)))
		return
	}
	helper104_0(root.Left, depth + 1)
	helper104_0(root.Right, depth + 1)
}
// 自底向上的递归
func maxDepth104_1(root *TreeNode) int {
	if root == nil { return 0 }
	left := maxDepth104_1(root.Left) + 1
	right := maxDepth104_1(root.Right) + 1
	return int(math.Max(float64(left), float64(right)))
}
// 与上面的形式略有差别
func maxDepth104_2(root *TreeNode) int {
	if root == nil { return 0 }
	left := maxDepth104_2(root.Left)
	right := maxDepth104_2(root.Right)
	return int(math.Max(float64(left), float64(right))) + 1
}