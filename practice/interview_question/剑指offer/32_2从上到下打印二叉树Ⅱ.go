package offer

// bfs
func levelOrder(root *TreeNode) [][]int {
	if root == nil { return nil }
	var ans [][]int
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		level := []int{}
		for i := 0; i < length; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			level = append(level, curr.Val)
		}
		ans = append(ans, level)
	}
	return ans
}
