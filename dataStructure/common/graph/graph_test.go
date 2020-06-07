package graph

import (
	"testing"
	"fmt"
)

func TestGraph(t *testing.T) {
	/*

	Graph:
	0 -- 1 -- 2
	|    |    |
	3 -- 4 -- 5
	     |    |
	     6 -- 7

	 */
	g := NewGraph(8)
	g.AddEdge(0, 1)
	g.AddEdge(0, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(4, 6)
	g.AddEdge(4, 5)
	g.AddEdge(5, 7)
	g.AddEdge(6, 7)
	fmt.Println("BFS:", g.BFS(2, 7))
	fmt.Println("DFS:", g.DFS(2, 7))
	// Output:
	// BFS: [2 5 7]
	// DFS: [2 1 0 3 4 6 7]
}
