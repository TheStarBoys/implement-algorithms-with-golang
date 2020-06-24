package leetcode

func recoverFromPreorder(S string) *TreeNode {
	path, pos := []*TreeNode{}, 0
	for pos < len(S) {
		level, value := 0, 0
		for pos < len(S) && S[pos] == '-' {
			pos++
			level++
		}
		for pos < len(S) && S[pos] != '-' {
			value = value * 10 + int(S[pos] - '0')
			pos++
		}
		node := &TreeNode{Val: value}
		if level == len(path) {
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			path = path[:level]
			path[len(path)-1].Right = node
		}
		path = append(path, node)
	}

	return path[0]
}
