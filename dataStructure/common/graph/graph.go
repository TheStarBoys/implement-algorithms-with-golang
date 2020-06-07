package graph

import (
	"container/list"
	"fmt"
)

// Graph 无向图
type Graph struct {
	v int // 顶点个数
	adj []list.List // 邻接表
}

func NewGraph(v int) *Graph {
	return &Graph{
		v: v,
		adj: make([]list.List, v),
	}
}

// AddEdge 无向图一条边存两次
func (g *Graph) AddEdge(s, t int) {
	if !(0 <= s && s < len(g.adj)) ||
		!(0 <= t && t < len(g.adj)) {
		// index out of range
		return
	}
	// 注意：没有判断边可能会重复
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
	fmt.Printf("adj[%d]: ", s)
	printList(g.adj[s])
	fmt.Printf("adj[%d]: ", t)
	printList(g.adj[t])
}

// bfs 求从 s 到 t 的最短路径，最短路径可能有多个，返回其中一个结果。
// 最坏情况下，终止顶点 t 离起始顶点 s 很远，需要遍历完整个图才能找到。
// 这个时候，每个顶点都要进出一遍队列，每个边也都会被访问一次，所以，广度优先搜索的时间复杂度是 O(V+E)
// 其中，V 表示顶点的个数，E 表示边的个数。当然，对于一个连通图来说，也就是说一个图中的所有顶点都是连通的
// E 肯定要大于等于 V-1，所以，广度优先搜索的时间复杂度也可以简写为 O(E)。

// 广度优先搜索的空间消耗主要在几个辅助变量 visited 数组、queue 队列、prev 数组上。
// 这三个存储空间的大小都不会超过顶点的个数，所以空间复杂度是 O(V)。
func (g Graph) BFS(s, t int) []int {
	if !(0 <= s && s < len(g.adj)) ||
		!(0 <= t && t < len(g.adj)) {
		// index out of range
		return nil
	}
	if s == t { return []int{s} }
	visited := make([]bool, g.v)
	visited[s] = true
	queue := make([]int, 0, g.v)
	queue = append(queue, s)
	// 存储从 顶点 t 到 顶点 s 的反向路径
	// 例如顶点 1 到顶点2，那么 prev[2] == 1
	prev := make([]int, g.v)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			l := g.adj[cur]
			for e := l.Front(); e != nil; e = e.Next() {
				v := e.Value.(int)
				if !visited[v] {
					prev[v] = cur
					if v == t {
						return reverse2(prev, s, t)
					}
					queue = append(queue, v)
					visited[v] = true
				}
			}
		}
	}

	// 不存在该路径
	return nil
}

// 每条边最多会被访问两次，一次是遍历，一次是回退。所以，图上的深度优先搜索算法的时间复杂度是 O(E)，E 表示边的个数。

// 深度优先搜索算法的消耗内存主要是 visited、path 数组和递归调用栈。visited、path 切片的大小跟顶点的个数 V 成正比，
// 递归调用栈的最大深度不会超过顶点的个数，所以总的空间复杂度就是 O(V)。
func (g *Graph) DFS(s, t int) []int {
	if !(0 <= s && s < len(g.adj)) ||
		!(0 <= t && t < len(g.adj)) {
		// index out of range
		return nil
	}
	visited := make([]bool, g.v)
	var path []int
	if ok, res := g.dfs(visited, path, s, t); ok {
		return res
	}

	return nil
}

func (g *Graph) dfs(visited []bool, path []int, s, t int) (bool, []int) {
	if s == t { return true, append(path, s) }
	visited[s] = true
	path = append(path, s)
	for e := g.adj[s].Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		if !visited[v] {
			ok, res := g.dfs(visited, path, v, t)
			if ok {
				return true, res
			}
		}
	}

	return false, nil
}

// 将反向路径转为正向路径
func reverse(prev []int, s, t int) []int {
	// 根据prev获取反向路径
	res := []int{t}
	for t != s {
		res = append(res, prev[t])
		t = prev[t]
	}
	// 反转res，得到其正向路径
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return res
}

// 将反向路径转为正向路径，用到了递归
func reverse2(prev []int, s, t int) []int {
	if s == t { return []int{s} }
	return append(reverse2(prev, s, prev[t]), t)
}

func printList(list list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}