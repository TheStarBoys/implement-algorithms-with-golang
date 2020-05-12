# 广度优先搜索

## 简介

广度优先搜索的预备知识——[队列](../dataStructure/stackAndQueue.md)

LeetCode官方简介内容——[队列和BFS](https://leetcode-cn.com/explore/learn/card/queue-stack/217/queue-and-bfs/869/)

## 模板Ⅰ

**Java**

```java
/**
 * Return the length of the shortest path between root and target node.
 */
int BFS(Node root, Node target) {
    Queue<Node> queue;  // store all nodes which are waiting to be processed
    int step = 0;       // number of steps neeeded from root to current node
    // initialize
    add root to queue;
    // BFS
    while (queue is not empty) {
        step = step + 1;
        // iterate the nodes which are already in the queue
        int size = queue.size();
        for (int i = 0; i < size; ++i) {
            Node cur = the first node in queue;
            return step if cur is target;
            for (Node next : the neighbors of cur) {
                add next to queue;
            }
            remove the first node from queue;
        }
    }
    return -1;          // there is no path from root to target
}
```

1. 如代码所示，在每一轮中，队列中的结点是`等待处理的结点`。
2. 在每个更外一层的 `while` 循环之后，我们`距离根结点更远一步`。变量 `step` 指示从根结点到我们正在访问的当前结点的距离。

## 模板Ⅱ

 有时，确保我们永远`不会访问一个结点两次`很重要。否则，我们可能陷入无限循环。如果是这样，我们可以在上面的代码中添加一个哈希集来解决这个问题。这是修改后的伪代码： 

**Java**

```java
/**
 * Return the length of the shortest path between root and target node.
 */
int BFS(Node root, Node target) {
    Queue<Node> queue;  // store all nodes which are waiting to be processed
    Set<Node> used;     // store all the used nodes
    int step = 0;       // number of steps neeeded from root to current node
    // initialize
    add root to queue;
    add root to used;
    // BFS
    while (queue is not empty) {
        step = step + 1;
        // iterate the nodes which are already in the queue
        int size = queue.size();
        for (int i = 0; i < size; ++i) {
            Node cur = the first node in queue;
            return step if cur is target;
            for (Node next : the neighbors of cur) {
                if (next is not in used) {
                    add next to queue;
                    add next to used;
                }
            }
            remove the first node from queue;
        }
    }
    return -1;          // there is no path from root to target
}
```

有两种情况你不需要使用哈希集：

1. 你完全确定没有循环，例如，在树遍历中；
2. 你确实希望多次将结点添加到队列中。

## 题目实战

### 1. 岛屿数量

#### 题目描述

给你一个由 `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

 

**示例 1:**

```
输入:
11110
11010
11000
00000
输出: 1
```

#### 题目分析

当遇到网格中值为`'1'`的时候，说明找到了一个岛屿，并且横竖相邻的均为该岛屿的土地。我们需要消除这些干扰因素。当找到一块岛屿的时候，通过广度优先搜索，将连在这块岛屿上的所有的`'1'`都清除掉，那么在清除掉后，再遇到的`'1'`，一定是一块新的岛屿。

#### 代码实现

**Go**

```go
func numIslands(grid [][]byte) int {
    row = len(grid)
	if row == 0 {
		return 0
	}

	col = len(grid[0])
    var count int
    
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            // 说明找到了一块岛屿
            if grid[i][j] == '1' {
                count++
                bfs(grid, i, j)
            }
        }
    }
    
    return count
}
// 上 [i-1, j] 下 [i+1, j] 左 [i, j-1] 右 [i, j+1]
// 偏移量，用于确定相对于当前位置的 下一个搜索目标
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}
var row, col int
// 用于清除'1'
func bfs(grid [][]byte, i, j int) {
    queue := make([]int, 0)
    // 队列中每两个元素，对应为坐标i, j
    queue = append(queue, i, j)
    grid[i][j] = '0'
    
    for len(queue) > 0 {
        // 取出当前队列中最靠前的坐标
        i, j = queue[0], queue[1]
        queue = queue[2:]
        // 确定下一个搜寻的下标
        for m := 0; m < 4; m++ {
            tmp_i := i + dx[m]
            tmp_j := j + dy[m]
            // 确保搜寻目标在网格内，并且是岛屿的一部分
            if 0 <= tmp_i && tmp_i < row && 0 <= tmp_j && tmp_j < col && grid[tmp_i][tmp_j] == '1' {
                grid[tmp_i][tmp_j] = '0'
                // 将待搜寻目标添加至队列
                queue = append(queue, tmp_i, tmp_j)
            }
        }
    }
}
```



### 2. 打开转盘锁

#### 题目描述

你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： `'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'` 。每个拨轮可以自由旋转：例如把 `'9'` 变为 `'0'`，`'0'` 变为 `'9'` 。每次旋转都只能旋转一个拨轮的一位数字。

锁的初始数字为 `'0000'` ，一个代表四个拨轮的数字的字符串。

列表 `deadends` 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

字符串 `target` 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。

 

**示例 1:**

```
输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。
```

#### 题目分析

首先判断题目，需要确定解锁的最小旋转次数，显然需要用bfs来做
由于每次都需要判断 你做出的旋转 得到的数字是否在死亡列表，因此用一个set来存储该列表，减少时间复杂度
确定如何选择下一次旋转：有四个拨轮，首先需要确定旋转哪一个拨轮，然后再确定是往上旋转还是往下旋转
记住 数字的取值范围在 0 ~ 9
意味着，如果 8 往上旋转为 9 ，9 再往上旋转 将回到 0
旋转拨轮的小技巧： '3' - '0' ==> 3, 3 + 1 ==> 4, 4 + '0' ==> '4'

并且需要确保所做的选择没有在之前做过

#### 代码实现

**Go**

```go
func openLock(deadends []string, target string) int {
	set := make(map[string]bool)
	for _, d := range deadends {
		set[d] = true
	}
	// 对特殊情况做处理
	if set[target] || set["0000"] {
		return -1
	}
	// 拨轮
	wheel := []byte{'0', '0', '0', '0'}
	return bfs(wheel, target, set)
}

func bfs(wheel []byte, target string, deadSet map[string]bool) int {
	step := 0
	queue := make([][]byte, 0)
	queue = append(queue, wheel)
	used := make(map[string]bool)
	used["0000"] = true

	for len(queue) > 0 {
		length := len(queue)
		step++
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			var nexts []string
			// 确定旋转哪一个拨轮
			for j := 0; j < 4; j++ {
				// 确定转动的方向
				newWheel := append([]byte{}, cur...)
				o := newWheel[j] // 保存该值，用于还原
				// 向上转
				newWheel[j] = (newWheel[j] - '0' + 1) % 10 +'0'
				nexts = append(nexts, string(newWheel))
				newWheel[j] = o // 撤销改动
				// 向下转
				newWheel[j] = (newWheel[j] - '0' + 9) % 10 + '0'
				nexts = append(nexts, string(newWheel))
				newWheel[j] = o // 撤销改动
			}
			// 从下一步操作的候选列表中，选出符合条件的
			for _, next := range nexts {
				if target == next {
					return step
				}
				if !deadSet[next] && !used[next] {
					queue = append(queue, []byte(next))
					used[next] = true
				}
			}
		}
	}

	return -1
}
```

