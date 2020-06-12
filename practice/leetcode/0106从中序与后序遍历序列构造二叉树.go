package leetcode

// 参考题解：
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/hou-xu-bian-li-python-dai-ma-java-dai-ma-by-liwe-2/
func buildTree106_0(inorder []int, postorder []int) *TreeNode {
	inlen, postlen := len(inorder), len(postorder)
	if inlen == 0 { return nil }
	if inlen == 1 { return &TreeNode{inorder[0], nil, nil} }
	rootVal := postorder[postlen - 1]
	dividePoint := 0
	for i := 0; i < inlen; i++ {
		if inorder[i] == rootVal {
			dividePoint = i
			break
		}
	}
	rootNode := &TreeNode{rootVal, nil, nil}
	rootNode.Left = buildTree106_0(inorder[0: dividePoint], postorder[0: dividePoint])
	rootNode.Right = buildTree106_0(inorder[dividePoint + 1:], postorder[dividePoint: postlen - 1])
	return rootNode
}
