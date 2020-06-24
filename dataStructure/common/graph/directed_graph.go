package graph

import (
	"container/list"
	"fmt"
)

type DiGraph struct {
	v int // 顶点个数
	adj []list.List
}

func NewDirectedGraph(v int) *DiGraph {
	return &DiGraph{
		v: v,
		adj: make([]list.List, v),
	}
}

func (g *DiGraph) AddEdge(s, t int) { // s 先于 t，边 s -> t
	if !(0 <= s && s < g.v) ||
		!(0 <= t && t < g.v) {
		// index out of range
		return
	}
	g.adj[s].PushBack(t)
}

func (g *DiGraph) TopoSortByKahn() {
	inDegree := make([]int, g.v) // 统计每个顶点的入度
	for i := 0; i < g.v; i++ {
		for e := g.adj[i].Front(); e != nil; e = e.Next() {
			val := e.Value.(int)
			inDegree[val]++
		}
	}
	queue := []int{}
	for i := 0; i < g.v; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		fmt.Printf("->%d", i)
		for e := g.adj[i].Front(); e != nil; e = e.Next() {
			val := e.Value.(int)
			inDegree[val]--
			if inDegree[val] == 0 {
				queue = append(queue, val)
			}
		}
	}
	fmt.Println()
}

func (g *DiGraph) TopoSortByDFS() {
	// 先构建逆邻接表，边 s -> t 表示， s 依赖于 t，t 先于 s
	inverseAdj := make([]list.List, g.v)
	for i := 0; i < g.v; i++ { // 通过邻接表生成逆邻接表
		for e := g.adj[i].Front(); e != nil; e = e.Next() {
			val := e.Value.(int)
			inverseAdj[val].PushBack(i)
		}
	}
	visited := make([]bool, g.v)
	for i := 0; i < g.v; i++ {
		if visited[i] { continue }
		g.dfs(i, inverseAdj, visited)
	}
	fmt.Println()
}

func (g *DiGraph) dfs(vertex int, inverseAdj []list.List, visited []bool) {
	visited[vertex] = true
	for e := inverseAdj[vertex].Front(); e != nil; e = e.Next() {
		val := e.Value.(int)
		if visited[val] == false {
			g.dfs(val, inverseAdj, visited)
		}
	}
	// 先把 vertex 这个顶点可达的所有顶点都打印出来之后，再打印它自己
	fmt.Printf("->%d", vertex)
}