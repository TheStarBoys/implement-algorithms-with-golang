# 深度优先搜索

## 简介

深度优先搜索的预备知识——[栈](../dataStructure/stackAndQueue.md)

LeetCode官方简介内容——[栈和BFS](https://leetcode-cn.com/explore/learn/card/queue-stack/219/stack-and-dfs/881/)

## 模板Ⅰ

正如我们在本章的描述中提到的，在大多数情况下，我们在能使用 BFS 时也可以使用 DFS。但是有一个重要的区别：`遍历顺序`。

与 BFS 不同，`更早访问的结点可能不是更靠近根结点的结点`。因此，你在 DFS 中找到的第一条路径`可能不是最短路径`。

在本文中，我们将为你提供一个 DFS 的递归模板，并向你展示栈是如何帮助这个过程的。在这篇文章之后，我们会提供一些练习给大家练习。

**Java**

```java
/*
 * Return true if there is a path from cur to target.
 */
boolean DFS(Node cur, Node target, Set<Node> visited) {
    return true if cur is target;
    for (next : each neighbor of cur) {
        if (next is not in visited) {
            add next to visted;
            return true if DFS(next, target, visited) == true;
        }
    }
    return false;
}
```

 当我们递归地实现 DFS 时，似乎不需要使用任何栈。但实际上，我们使用的是由系统提供的隐式栈，也称为调用栈（[Call Stack](https://en.wikipedia.org/wiki/Call_stack)）。 

## 模版Ⅱ

递归解决方案的优点是它更容易实现。 但是，存在一个很大的缺点：如果递归的深度太高，你将遭受堆栈溢出。 在这种情况下，您可能会希望使用 BFS，或使用显式栈实现 DFS。

这里我们提供了一个使用显式栈的模板：

**Java**

```java
/*
 * Return true if there is a path from cur to target.
 */
boolean DFS(int root, int target) {
    Set<Node> visited;
    Stack<Node> s;
    add root to s;
    while (s is not empty) {
        Node cur = the top element in s;
        return true if cur is target;
        for (Node next : the neighbors of cur) {
            if (next is not in visited) {
                add next to s;
                add next to visited;
            }
        }
        remove cur from s;
    }
    return false;
}
```

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

#### 代码实现

```go
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    var count int
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                count++
                dfs(grid, i, j)
            }
        }
    }
    
    return count
}

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}
func dfs(grid [][]byte, i, j int) {
    grid[i][j] = '0'
    for k := 0; k < 4; k++ {
        tmp_i := i + dx[k]
        tmp_j := j + dy[k]
        if 0 <= tmp_i && tmp_i < len(grid) && 0 <= tmp_j && tmp_j < len(grid[0]) && grid[tmp_i][tmp_j] == '1' {
            dfs(grid, tmp_i, tmp_j)
        }
    }
}
```

