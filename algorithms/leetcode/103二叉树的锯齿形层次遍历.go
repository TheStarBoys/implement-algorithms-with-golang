package leetcode

// 就是层次遍历的变形，在输出数组的时候，判断要不要反转就行了
func zigzagLevelOrder103_0(root *TreeNode) [][]int {
	if root == nil { return [][]int{} } // bfs要注意边界
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	ans := make([][]int, 0)
	flag := 0 // 1 反转

	for len(queue) != 0 {
		length := len(queue)
		ans = append(ans, []int{})
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			ans[depth] = append(ans[depth], cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		if flag == 1 {
			reverse103_0(ans[depth])
			flag = 0
		} else {
			flag = 1
		}
		depth++
	}
	return ans
}
func reverse103_0(nums []int) {
	l, r := 0, len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
