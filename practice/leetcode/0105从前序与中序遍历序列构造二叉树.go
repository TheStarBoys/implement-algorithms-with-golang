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


func buildTree105_1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0  {
		return nil
	}
	mid := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == mid.Val {
			break
		}
	}
	mid.Left = buildTree105_1(preorder[1 : i+1], inorder[:i])
	mid.Right = buildTree105_1(preorder[i+1 :], inorder[i+1:])

	return mid
}