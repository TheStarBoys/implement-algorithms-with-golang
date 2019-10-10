package main

import "fmt"

func main() {
	r1 := &TreeNode{3, &TreeNode{4,nil,nil}, &TreeNode{6, nil,nil}}
	r2 := &TreeNode{2, &TreeNode{5,nil,nil}, nil}
	root := &TreeNode{1, r2, r1}
	fmt.Println(postorderTraversal(root))
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	visited := make(map[*TreeNode]bool)
	visited[root] = true
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		node := cur.Left
		for node != nil && !visited[node] {
			visited[node] = true
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		if (node.Left == nil || visited[node.Left]) &&
			(node.Right == nil || visited[node.Right]) {
			ans = append(ans, node.Val)
			stack = stack[:len(stack)-1] // 注意这里的坑，在输出值的时候，才出栈
		}
		if node.Right != nil && !visited[node.Right] {
			stack = append(stack, node.Right)
			visited[node.Right] = true
		}

	}
	return ans
}