package leetcode

func buildTree105_0(preorder []int, inorder []int) *TreeNode {
	prelen, inlen := len(preorder), len(inorder)
	if prelen == 0 { return nil }
	if prelen == 1 { return &TreeNode{preorder[0], nil, nil} }

	rootVal := preorder[0]
	dividePoint := 0
	for i := 0; i < inlen; i++ {
		if inorder[i] == rootVal {
			dividePoint = i
			break
		}
	}
	rootNode := &TreeNode{rootVal, nil, nil}
	rootNode.Left = buildTree105_0(preorder[1: dividePoint + 1], inorder[:dividePoint])
	rootNode.Right = buildTree105_0(preorder[dividePoint + 1:], inorder[dividePoint + 1:])

	return rootNode
}
