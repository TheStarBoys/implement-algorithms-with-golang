package leetcode

/*

type Node struct {
	Val int
	Neighbors []*Node
}


func cloneGraph(node *Node) *Node {
	set := make(map[*Node]*Node)
	return dfs(node, set)
}

func dfs(cur *Node, set map[*Node]*Node) *Node {
	if cur == nil {
		return nil
	}
	if set[cur] != nil {
		return set[cur]
	}
	var cloneNode Node
	cloneNode.Val = cur.Val
	set[cur] = &cloneNode
	for i := 0; i < len(cur.Neighbors); i++ {
		cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(cur.Neighbors[i], set))
	}

	return &cloneNode
}

*/