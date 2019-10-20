package leetcode

// 先将二叉树中序遍历，得到序列，再验证是否是升序序列，来判断是否是BST
func isValidBST098_0(root *TreeNode) bool {
	list := []int{}
	list = getNums098_0(list, root)
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false
		}
	}
	return true
}

func getNums098_0(ans []int, root *TreeNode) []int {
	if root == nil { return ans }
	ans = getNums098_0(ans, root.Left)
	ans = append(ans, root.Val)
	ans = getNums098_0(ans, root.Right)
	return ans
}

// 在中序遍历的时候，就完成了验证
func isValidBST098_1(root *TreeNode) bool {
	list := []int{}
	list, set := getNums098_1(list, root)
	return set
}

func getNums098_1(ans []int, root *TreeNode) ([]int, bool) {
	if root == nil { return ans, true }
	var ok bool // 注意，if里面不能用:=，否则ans的值接收不到
	if ans, ok = getNums098_1(ans, root.Left); !ok { return ans, ok }
	if len(ans) != 0 && root.Val <= ans[len(ans)-1] { return ans, false }
	ans = append(ans, root.Val)
	if ans, ok = getNums098_1(ans, root.Right); !ok { return ans, ok }
	return ans, true
}