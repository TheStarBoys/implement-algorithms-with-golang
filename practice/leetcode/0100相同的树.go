package leetcode

// 100
func isSameTree100_0(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { return true }
	if p != nil && q == nil || p == nil && q != nil ||
		p.Val != q.Val {
		return false
	}
	return isSameTree100_0(p.Left, q.Left) && isSameTree100_0(p.Right, q.Right)
}
