package leetcode

// 递归
var ans *TreeNode
func lowestCommonAncestor236_0(root, p, q *TreeNode) *TreeNode {
	recurseTree236_0(root, p, q)
	return ans
}

func recurseTree236_0(currentNode, p, q *TreeNode) bool {
	if currentNode == nil { return false }
	left := 0
	if recurseTree236_0(currentNode.Left, p, q) {
		left = 1
	}
	right := 0
	if recurseTree236_0(currentNode.Right, p, q) {
		right = 1
	}
	mid := 0
	if currentNode == p || currentNode == q {
		mid = 1
	}
	if mid + left + right >= 2 {
		ans = currentNode
	}
	return mid + left + right > 0
}