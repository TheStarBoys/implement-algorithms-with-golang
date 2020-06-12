package leetcode

 //Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return make([]*TreeNode, 0)
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if start > end {
		trees = append(trees, nil)
		return trees
	}
	var root *TreeNode
	var left, right []*TreeNode
	for i := start; i <= end; i++ {
		left = generate(start, i-1)
		right = generate(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root = &TreeNode{i, nil, nil}
				root.Left = l
				root.Right = r
				trees = append(trees, root)
			}
		}
	}
	return trees
}

