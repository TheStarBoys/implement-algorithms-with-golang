# N叉树

## 简介

> 引用自：[LeetCode](https://leetcode-cn.com/explore/learn/card/n-ary-tree/)
>
> 代码实现：TheStarBoys

此章至少需要的预备知识：

- [二叉树](./binaryTree.md)



## N叉树的定义

二叉树是一棵以根节点开始，每个节点含有不超过2个子节点的树。 让我们将这个定义扩展到`N叉树`。 一棵以根节点开始，每个节点不超过`N`个子节点的树，称为`N叉树`。

以下是三叉树的一个例子：



![img](https://s3-us-west-1.amazonaws.com/s3-lc-upload/explore/cards/n-ary-tree/nary_tree_example.png)





`前缀树`，又称`字典树`(Trie)， 就是一个常用的N叉树。

并且二叉树是N叉树的一种特殊表示形式。在接下来的章节中，我们将把你所掌握的与二叉树相关的知识全部拓展到N叉树中。

## N叉树的数据结构定义

```go
type Node struct {
	Val int
	Children []*Node
}
```



## N叉树的遍历

### 树的遍历

------

一棵二叉树可以按照前序、中序、后序或者层序来进行遍历。在这些遍历方法中，前序遍历、后序遍历和层序遍历同样可以运用到N叉树中。

> 回顾 - 二叉树的遍历
>
> 1. 前序遍历 - 首先访问根节点，然后遍历左子树，最后遍历右子树；
> 2. 中序遍历 - 首先遍历左子树，然后访问根节点，最后遍历右子树；
> 3. 后序遍历 - 首先遍历左子树，然后遍历右子树，最后访问根节点；
> 4. 层序遍历 - 按照从左到右的顺序，逐层遍历各个节点。

请注意，N叉树的中序遍历没有标准定义，中序遍历只有在二叉树中有明确的定义。尽管我们可以通过几种不同的方法来定义N叉树的中序遍历，但是这些描述都不是特别贴切，并且在实践中也不常用到，所以我们暂且跳过N叉树中序遍历的部分。

把上述关于二叉树遍历转换为N叉树遍历，我们只需把如下表述:

> 遍历左子树... 遍历右子树... 

变为:

> 对于每个子节点:
>    通过递归地调用遍历函数来遍历以该子节点为根的子树

我们假设for循环将会按照各个节点在数据结构中的顺序进行遍历：通常按照从左到右的顺序，如下所示。

 

### N叉树遍历示例

------

我们用如图所示的三叉树来举例说明:

![img](https://s3-us-west-1.amazonaws.com/s3-lc-upload/explore/cards/n-ary-tree/nary_tree_example.png)

 

#### 1.前序遍历

在N叉树中，前序遍历指先访问根节点，然后逐个遍历以其子节点为根的子树。
例如，上述三叉树的前序遍历是: A->B->C->E->F->D->G.

#### 2.后序遍历

在N叉树中，后序遍历指前先逐个遍历以根节点的子节点为根的子树，最后访问根节点。
例如，上述三叉树的后序遍历是: B->E->F->C->G->D->A.

#### 3.层序遍历

N叉树的层序遍历与二叉树的一致。通常，当我们在树中进行广度优先搜索时，我们将按层序的顺序进行遍历。
例如，上述三叉树的层序遍历是: A->B->C->D->E->F->G.

## 题目实战

### [前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)

#### 题目描述

给定一个 N 叉树，返回其节点值的*前序遍历*。

例如，给定一个 `3叉树` :

 

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png)

 

返回其前序遍历: `[1,3,5,6,2,4]`。

 

**说明:** 递归法很简单，你可以使用迭代法完成此题吗?

#### 代码实现

**方法一：递归**

```go
func preorder(root *Node) []int {
    var ans []int
    helper(root, &ans)
    
    return ans
}

func helper(root *Node, nums *[]int) {
    if root == nil {
        return
    }
    *nums = append(*nums, root.Val)
    for _, node := range root.Children {
        helper(node, nums)
    }
}
```

**方法二：迭代**

```go
func preorder(root *Node) []int {
    if root == nil { return nil }

    res := []int{}
    stack := []*Node{root}
    for len(stack) != 0 {
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, curr.Val)
        // 逆序遍历 Children 入栈，来保证出栈是顺序的
        for i := len(curr.Children) - 1; i >= 0; i-- {
            child := curr.Children[i]
            stack = append(stack, child)
        }
    }
    return res
}
```



### [后序遍历](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)

#### 题目描述

给定一个 N 叉树，返回其节点值的*后序遍历*。

例如，给定一个 `3叉树` :

 

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png)

 

返回其后序遍历: `[5,6,3,2,4,1]`.

 

**说明:** 递归法很简单，你可以使用迭代法完成此题吗?

#### 代码实现

**方法一：递归**

```go
func postorder(root *Node) []int {
    var ans []int
    helper(root, &ans)
    return ans
}

func helper(root *Node, nums *[]int) {
    if root == nil {
        return
    }
    for _, node := range root.Children {
        helper(node, nums)
    }
    *nums = append(*nums, root.Val)
}
```

**方法二：迭代**

```go
func postorder(root *Node) []int {
    if root == nil { return nil }

    res := []int{}
    stack := []*Node{root}
    visited := make(map[*Node]bool)
    for len(stack) != 0 {
        curr := stack[len(stack)-1]
        if visited[curr] {
            res = append(res, curr.Val)
            stack = stack[:len(stack)-1]
            continue
        }
        visited[curr] = true
        for i := len(curr.Children) - 1; i >= 0; i-- {
            child := curr.Children[i]
            stack = append(stack, child)
        }
    }

    return res
}
```



### [层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)

#### 题目描述

给定一个 N 叉树，返回其节点值的*层序遍历*。 (即从左到右，逐层遍历)。

例如，给定一个 `3叉树` :

 

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png)

 

返回其层序遍历:

```
[
     [1],
     [3,2,4],
     [5,6]
]
```

 

**说明:**

1. 树的深度不会超过 `1000`。
2. 树的节点总数不会超过 `5000`。

#### 代码实现

```go
func levelOrder(root *Node) [][]int {
    if root == nil { return nil }

    res := [][]int{}
    queue := []*Node{root}
    
    for level := 0; len(queue) != 0; level++ {
        length := len(queue)
        res = append(res, []int{})
        for i := 0; i < length; i++ {
            curr := queue[0]
            queue = queue[1:]
            res[level] = append(res[level], curr.Val)
            for _, child := range curr.Children {
                queue = append(queue, child)
            }
        }
    }

    return res
}
```

### [N叉树最大深度](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/)

#### 题目描述

给定一个 N 叉树，找到其最大深度。

最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

例如，给定一个 `3叉树` :

 

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/narytreeexample.png)

 

我们应返回其最大深度，3。

**说明:**

1. 树的深度不会超过 `1000`。
2. 树的节点总不会超过 `5000`。

#### 代码实现

```go
func maxDepth(root *Node) int {
    if root == nil { return 0 }
    if len(root.Children) == 0 { return 1 }
    max := math.MinInt64
    for _, child := range root.Children {
        if depth := maxDepth(child); depth > max {
            max = depth
        }
    }

    return max + 1
}
```

