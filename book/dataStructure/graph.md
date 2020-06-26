# 图

## 简介

**图（Graph）**，和树比起来，这是一种更加复杂的非线性表结构。

我们知道，树中的元素我们称为节点，图中的元素我们就叫作**顶点（vertex）**。从我画的图中可以看出来，图中的一个顶点可以与任意其他顶点建立连接关系。我们把这种建立的关系叫作**边（edge）**。 

 ![img](https://static001.geekbang.org/resource/image/df/af/df85dc345a9726cab0338e68982fd1af.jpg) 

我们就拿微信举例子吧。我们可以把每个用户看作一个顶点。如果两个用户之间互加好友，那就在两者之间建立一条边。所以，整个微信的好友关系就可以用一张图来表示。其中，每个用户有多少个好友，对应到图中，就叫作顶点的**度（degree）**，就是跟顶点相连接的边的条数。 

实际上，微博的社交关系跟微信还有点不一样，或者说更加复杂一点。微博允许单向关注，也就是说，用户 A 关注了用户 B，但用户 B 可以不关注用户 A。那我们如何用图来表示这种单向的社交关系呢？

我们可以把刚刚讲的图结构稍微改造一下，引入边的“方向”的概念。

如果用户 A 关注了用户 B，我们就在图中画一条从 A 到 B 的带箭头的边，来表示边的方向。如果用户 A 和用户 B 互相关注了，那我们就画一条从 A 指向 B 的边，再画一条从 B 指向 A 的边。我们把这种边有方向的图叫作“有向图”。以此类推，我们把边没有方向的图就叫作“无向图”。 

 ![img](https://static001.geekbang.org/resource/image/c3/96/c31759a37d8a8719841f347bd479b796.jpg) 

我们刚刚讲过，无向图中有“度”这个概念，表示一个顶点有多少条边。在有向图中，我们把度分为**入度（In-degree）**和**出度（Out-degree）**。 

顶点的入度，表示有多少条边指向这个顶点；顶点的出度，表示有多少条边是以这个顶点为起点指向其他顶点。对应到微博的例子，入度就表示有多少粉丝，出度就表示关注了多少人。 

前面讲到了微信、微博、无向图、有向图，现在我们再来看另一种社交软件：QQ。

QQ 中的社交关系要更复杂的一点。不知道你有没有留意过 QQ 亲密度这样一个功能。QQ 不仅记录了用户之间的好友关系，还记录了两个用户之间的亲密度，如果两个用户经常往来，那亲密度就比较高；如果不经常往来，亲密度就比较低。如何在图中记录这种好友关系的亲密度呢？

这里就要用到另一种图，**带权图（weighted graph）**。在带权图中，每条边都有一个权重（weight），我们可以通过这个权重来表示 QQ 好友间的亲密度。 

 ![img](https://static001.geekbang.org/resource/image/55/e8/55d7e4806dc47950ae098d959b03ace8.jpg) 

## 图的存储方式

### 邻接矩阵存储方法

图最直观的一种存储方法就是，**邻接矩阵（Adjacency Matrix）**。 

邻接矩阵的底层依赖一个二维数组。对于无向图来说，如果顶点 `i` 与顶点 `j` 之间有边，我们就将 `A[i][j]` 和 `A[j][i]` 标记为 1；对于有向图来说，如果顶点`i` 到顶点 `j` 之间，有一条箭头从顶点 `i` 指向顶点 `j` 的边，那我们就将 `A[i][j]` 标记为 1。同理，如果有一条箭头从顶点 `j` 指向顶点 `i` 的边，我们就将 `A[j][i]` 标记为 1。对于带权图，数组中就存储相应的权重。 

 ![img](https://static001.geekbang.org/resource/image/62/d2/625e7493b5470e774b5aa91fb4fdb9d2.jpg) 

用邻接矩阵来表示一个图，虽然简单、直观，但是比较浪费存储空间。为什么这么说呢？对于无向图来说，如果 `A[i][j]` 等于 1，那 `A[j][i]` 也肯定等于 1。实际上，我们只需要存储一个就可以了。也就是说，无向图的二维数组中，如果我们将其用对角线划分为上下两部分，那我们只需要利用上面或者下面这样一半的空间就足够了，另外一半白白浪费掉了。 

还有，如果我们存储的是**稀疏图（Sparse Matrix）**，也就是说，顶点很多，但每个顶点的边并不多，那邻接矩阵的存储方法就更加浪费空间了。比如微信有好几亿的用户，对应到图上就是好几亿的顶点。但是每个用户的好友并不会很多，一般也就三五百个而已。如果我们用邻接矩阵来存储，那绝大部分的存储空间都被浪费了。

但这也并不是说，邻接矩阵的存储方法就完全没有优点。首先，邻接矩阵的存储方式简单、直接，因为基于数组，所以在获取两个顶点的关系时，就非常高效。其次，用邻接矩阵存储图的另外一个好处是方便计算。这是因为，用邻接矩阵的方式存储图，可以将很多图的运算转换成矩阵之间的运算。比如求解最短路径问题时会提到一个[Floyd-Warshall 算法](https://zh.wikipedia.org/wiki/Floyd-Warshall算法) ，就是利用矩阵循环相乘若干次得到结果。 

### 邻接表存储方法

针对上面邻接矩阵比较浪费内存空间的问题，我们来看另外一种图的存储方法，**邻接表（Adjacency List）**。 

我画了一张邻接表的图，你可以先看下。乍一看，邻接表是不是有点像散列表？每个顶点对应一条链表，链表中存储的是与这个顶点相连接的其他顶点。另外我需要说明一下，图中画的是一个有向图的邻接表存储方式，每个顶点对应的链表里面，存储的是指向的顶点。对于无向图来说，也是类似的，不过，每个顶点的链表中存储的，是跟这个顶点有边相连的顶点，你可以自己画下。 

 ![img](https://static001.geekbang.org/resource/image/03/94/039bc254b97bd11670cdc4bf2a8e1394.jpg) 

还记得我们之前讲过的时间、空间复杂度互换的设计思想吗？邻接矩阵存储起来比较浪费空间，但是使用起来比较节省时间。相反，邻接表存储起来比较节省空间，但是使用起来就比较耗时间。 

就像图中的例子，如果我们要确定，是否存在一条从顶点 2 到顶点 4 的边，那我们就要遍历顶点 2 对应的那条链表，看链表中是否存在顶点 4。而且，我们前面也讲过，链表的存储方式对缓存不友好。所以，比起邻接矩阵的存储方式，在邻接表中查询两个顶点之间的关系就没那么高效了。 

在散列表那几节里，我讲到，在基于链表法解决冲突的散列表中，如果链过长，为了提高查找效率，我们可以将链表换成其他更加高效的数据结构，比如平衡二叉查找树等。我们刚刚也讲到，邻接表长得很像散列。所以，我们也可以将邻接表同散列表一样进行“改进升级”。 

我们可以将邻接表中的链表改成平衡二叉查找树。实际开发中，我们可以选择用红黑树。这样，我们就可以更加快速地查找两个顶点之间是否存在边了。当然，这里的二叉查找树可以换成其他动态数据结构，比如跳表、散列表等。除此之外，我们还可以将链表改成有序动态数组，可以通过二分查找的方法来快速定位两个顶点之间否是存在边。 

## 图的实现

### 邻接表实现图

#### 无向图

```go
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
}
```



##### 图的广度优先搜索

```go
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
						return reverse(prev, s, t)
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
```

掌握了广优先搜索算法的原理，我们来看下，广度优先搜索的时间、空间复杂度是多少呢？

最坏情况下，终止顶点 t 离起始顶点 s 很远，需要遍历完整个图才能找到。这个时候，每个顶点都要进出一遍队列，每个边也都会被访问一次，所以，广度优先搜索的时间复杂度是 `O(V+E)`，其中，V 表示顶点的个数，E 表示边的个数。当然，对于一个连通图来说，也就是说一个图中的所有顶点都是连通的，E 肯定要大于等于 V-1，所以，广度优先搜索的时间复杂度也可以简写为 `O(E)`。 

广度优先搜索的空间消耗主要在几个辅助变量 visited 数组、queue 队列、prev 数组上。这三个存储空间的大小都不会超过顶点的个数，所以空间复杂度是 `O(V)`。 

##### 图的深度优先搜索

```go
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
```

你可以看我画的这幅图。搜索的起始顶点是 s，终止顶点是 t，我们希望在图中寻找一条从顶点 s 到顶点 t 的路径。如果映射到迷宫那个例子，s 就是你起始所在的位置，t 就是出口。

我用深度递归算法，把整个搜索的路径标记出来了。这里面实线箭头表示遍历，虚线箭头表示回退。从图中我们可以看出，深度优先搜索找出来的路径，并不是顶点 s 到顶点 t 的最短路径。 ![img](https://static001.geekbang.org/resource/image/87/85/8778201ce6ff7037c0b3f26b83efba85.jpg) 

理解了深度优先搜索算法之后，我们来看，深度度优先搜索的时、空间间复杂度是多少呢？

从我前面画的图可以看出，每条边最多会被访问两次，一次是遍历，一次是回退。所以，图上的深度优先搜索算法的时间复杂度是 `O(E)`，E 表示边的个数。

深度优先搜索算法的消耗内存主要是 visited、path 数组和递归调用栈。visited、path 数组的大小跟顶点的个数 V 成正比，递归调用栈的最大深度不会超过顶点的个数，所以总的空间复杂度就是 `O(V)`。 

##### 示例

```go
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
```

#### 有向图

```go
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
```

##### 拓扑排序

那什么是拓扑排序呢？这个概念很好理解，我们先来看一个生活中的拓扑排序的例子。

我们在穿衣服的时候都有一定的顺序，我们可以把这种顺序想成，衣服与衣服之间有一定的依赖关系。比如说，你必须先穿袜子才能穿鞋，先穿内裤才能穿秋裤。

假设我们现在有八件衣服要穿，它们之间的两两依赖关系我们已经很清楚了，那如何安排一个穿衣序列，能够满足所有的两两之间的依赖关系？这就是个拓扑排序问题。从这个例子中，你应该能想到，在很多时候，拓扑排序的序列并不是唯一的。你可以看我画的这幅图，我找到了好几种满足这些局部先后关系的穿衣序列。 

 ![img](https://static001.geekbang.org/resource/image/c2/bd/c26d0f472d9a607c0c4eb688c01959bd.jpg) 

弄懂了这个生活中的例子，开篇的关于编译顺序的问题，你应该也有思路了。开篇问题跟这个问题的模型是一样的，也可以抽象成一个拓扑排序问题。

拓扑排序的原理非常简单，我们的重点应该放到拓扑排序的实现上面。我前面多次讲过，算法是构建在具体的数据结构之上的。针对这个问题，我们先来看下，如何将问题背景抽象成具体的数据结构？

我们可以把源文件与源文件之间的依赖关系，抽象成一个有向图。每个源文件对应图中的一个顶点，源文件之间的依赖关系就是顶点之间的边。如果 a 先于 b 执行，也就是说 b 依赖于 a，那么就在顶点 a 和顶点 b 之间，构建一条从 a 指向 b 的边。而且，这个图不仅要是有向图，还要是一个有向无环图，也就是不能存在像 a->b->c->a 这样的循环依赖关系。因为图中一旦出现环，拓扑排序就无法工作了。实际上，拓扑排序本身就是基于有向无环图的一个算法。 

**1.Kahn 算法**

Kahn 算法实际上用的是贪心算法思想，思路非常简单、好懂。

定义数据结构的时候，如果 s 需要先于 t 执行，那就添加一条 s 指向 t 的边。所以，如果某个顶点入度为 0， 也就表示，没有任何顶点必须先于这个顶点执行，那么这个顶点就可以执行了。

我们先从图中，找出一个入度为 0 的顶点，将其输出到拓扑排序的结果序列中（对应代码中就是把它打印出来），并且把这个顶点从图中删除（也就是把这个顶点可达的顶点的入度都减 1）。我们循环执行上面的过程，直到所有的顶点都被输出。最后输出的序列，就是满足局部依赖关系的拓扑排序。

我把 Kahn 算法用代码实现了一下，你可以结合着文字描述一块看下。不过，你应该能发现，这段代码实现更有技巧一些，并没有真正删除顶点的操作。代码中有详细的注释，你自己来看，我就不多解释了。 

```go
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
```

**2.DFS 算法**

图上的深度优先搜索我们前面已经讲过了，实际上，拓扑排序也可以用深度优先搜索来实现。不过这里的名字要稍微改下，更加确切的说法应该是深度优先遍历，遍历图中的所有顶点，而非只是搜索一个顶点到另一个顶点的路径。

关于这个算法的实现原理，我先把代码贴在下面，下面给你具体解释。 

```go
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
```

这个算法包含两个关键部分。

第一部分是**通过邻接表构造逆邻接表**。邻接表中，边 s->t 表示 s 先于 t 执行，也就是 t 要依赖 s。在逆邻接表中，边 s->t 表示 s 依赖于 t，s 后于 t 执行。为什么这么转化呢？这个跟我们这个算法的实现思想有关。

第二部分是这个算法的核心，也就是**递归处理每个顶点**。对于顶点 vertex 来说，我们先输出它可达的所有顶点，也就是说，先把它依赖的所有的顶点输出了，然后再输出自己。 

## 引用

> [数据结构与算法之美](https://time.geekbang.org/column/intro/100017301)