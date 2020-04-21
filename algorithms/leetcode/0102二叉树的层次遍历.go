package leetcode

func levelOrder102_0(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil { return ans }
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for floor := 0; len(queue) != 0; floor++ {
		ans = append(ans, []int{})
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			ans[floor] = append(ans[floor], cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return ans
}
