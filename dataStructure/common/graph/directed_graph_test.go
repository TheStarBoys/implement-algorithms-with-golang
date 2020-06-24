package graph

import "testing"

func TestDiGraph(t *testing.T) {
	// 依赖关系：0 -> 1, 1 -> 2, 3 -> 1
	/*
			0 --> 1 --> 2
			3 ----^
	 */
	g := NewDirectedGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(3, 1)

	g.TopoSortByKahn()

	g.TopoSortByDFS()
}
