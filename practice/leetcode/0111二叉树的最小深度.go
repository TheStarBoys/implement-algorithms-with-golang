package leetcode

import "math"


func minDepth0111(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getMin0111(root)
}

func getMin0111(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := getMin0111(root.Left)
	right := getMin0111(root.Right)
	return int(math.Min(float64(left), float64(right))) + 1
}
