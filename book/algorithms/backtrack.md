# 回溯算法

> 引用自：[回溯算法详解（修订版）]( https://www.cxyxiaowu.com/7103.html )

## 解题框架

废话不多说，直接上回溯算法框架。**解决一个回溯问题，实际上就是一个决策树的遍历过程**。你只需要思考 3 个问题：

**1、路径：**也就是已经做出的选择。

**2、选择列表：**也就是你当前可以做的选择。

**3、结束条件**：也就是到达决策树底层，无法再做选择的条件。

如果你不理解这三个词语的解释，没关系，我们后面会用「全排列」和「N 皇后问题」这两个经典的回溯算法问题来帮你理解这些词语是什么意思，现在你先留着印象。

代码方面，回溯算法的框架：

```python
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```

**其核心就是 for 循环里面的递归，在递归调用之前「做选择」，在递归调用之后「撤销选择」**，特别简单。

什么叫做选择和撤销选择呢，这个框架的底层原理是什么呢？下面我们就通过「全排列」这个问题来解开之前的疑惑，详细探究一下其中的奥妙！

## 题目实战

### 1. 全排列问题

#### 题目描述

给定一个**没有重复**数字的序列，返回其所有可能的全排列。

**示例:**

```
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```



#### 题目解析

我们在高中的时候就做过排列组合的数学题，我们也知道`n`个不重复的数，全排列共有 n! 个。

PS：**为了简单清晰起见，我们这次讨论的全排列问题不包含重复的数字**。

那么我们当时是怎么穷举全排列的呢？比方说给三个数`[1,2,3]`，你肯定不会无规律地乱穷举，一般是这样：

先固定第一位为 1，然后第二位可以是 2，那么第三位只能是 3；然后可以把第二位变成 3，第三位就只能是 2 了；然后就只能变化第一位，变成 2，然后再穷举后两位……

其实这就是回溯算法，我们高中无师自通就会用，或者有的同学直接画出如下这棵回溯树：

![回溯算法详解（修订版）](https://www.cxyxiaowu.com/wp-content/uploads/2019/11/1574989306-e84276bc4f9bd5e.jpeg)

只要从根遍历这棵树，记录路径上的数字，其实就是所有的全排列。**我们不妨把这棵树称为回溯算法的「决策树」**。

**为啥说这是决策树呢，因为你在每个节点上其实都在做决策**。比如说你站在下图的红色节点上：

![回溯算法详解（修订版）](https://www.cxyxiaowu.com/wp-content/uploads/2019/11/1574989306-e84276bc4f9bd5e-1.jpeg)

你现在就在做决策，可以选择 1 那条树枝，也可以选择 3 那条树枝。为啥只能在 1 和 3 之中选择呢？因为 2 这个树枝在你身后，这个选择你之前做过了，而全排列是不允许重复使用数字的。

**现在可以解答开头的几个名词：`[2]`就是「路径」，记录你已经做过的选择；`[1,3]`就是「选择列表」，表示你当前可以做出的选择；「结束条件」就是遍历到树的底层，在这里就是选择列表为空的时候**。

如果明白了这几个名词，**可以把「路径」和「选择列表」作为决策树上每个节点的属性**，比如下图列出了几个节点的属性：

![回溯算法详解（修订版）](https://www.cxyxiaowu.com/wp-content/uploads/2019/11/1574989306-10769e3a4ac18b7.jpeg)

**我们定义的`backtrack`函数其实就像一个指针，在这棵树上游走，同时要正确维护每个节点的属性，每当走到树的底层，其「路径」就是一个全排列**。

再进一步，如何遍历一棵树？这个应该不难吧。回忆一下之前 [学习数据结构的框架思维](http://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484520&idx=1&sn=2c6507c7f25c0fd29fd1d146ee3b067c&chksm=9bd7fa60aca073763785418d15ed03c9debdd93ca36f4828fa809990116b1e7536c3f68a7b71&scene=21#wechat_redirect) 写过，各种搜索问题其实都是树的遍历问题，而多叉树的遍历框架就是这样：

```java
void traverse(TreeNode root) {
    for (TreeNode child : root.childern)
        // 前序遍历需要的操作
        traverse(child);
        // 后序遍历需要的操作
}
```

而所谓的前序遍历和后序遍历，他们只是两个很有用的时间点，我给你画张图你就明白了：

![回溯算法详解（修订版）](https://www.cxyxiaowu.com/wp-content/uploads/2019/11/1574989307-6ea51c4bd4b2c04.jpeg)

**前序遍历的代码在进入某一个节点之前的那个时间点执行，后序遍历代码在离开某个节点之后的那个时间点执行**。

回想我们刚才说的，「路径」和「选择」是每个节点的属性，函数在树上游走要正确维护节点的属性，那么就要在这两个特殊时间点搞点动作：

![回溯算法详解（修订版）](https://www.cxyxiaowu.com/wp-content/uploads/2019/11/1574989307-0e12d2bfc84e633.jpeg)

现在，你是否理解了回溯算法的这段核心框架？

```python
for 选择 in 选择列表:
    # 做选择
    将该选择从选择列表移除
    路径.add(选择)
    backtrack(路径, 选择列表)
    # 撤销选择
    路径.remove(选择)
    将该选择再加入选择列表
```

**我们只要在递归之前做出选择，在递归之后撤销刚才的选择**，就能正确得到每个节点的选择列表和路径。

#### 代码实现

**Java**

```java
List<List<Integer>> res = new LinkedList<>();

/* 主函数，输入一组不重复的数字，返回它们的全排列 */
List<List<Integer>> permute(int[] nums) {
    // 记录「路径」
    LinkedList<Integer> track = new LinkedList<>();
    backtrack(nums, track);
    return res;
}

// 路径：记录在 track 中// 选择列表：nums 中不存在于 track 的那些元素// 结束条件：nums 中的元素全都在 track 中出现void backtrack(int[] nums, LinkedList<Integer> track) {
    // 触发结束条件
    if (track.size() == nums.length) {
        res.add(new LinkedList(track));
        return;
    }

    for (int i = 0; i < nums.length; i++) {
        // 排除不合法的选择
        if (track.contains(nums[i]))
            continue;
        // 做选择
        track.add(nums[i]);
        // 进入下一层决策树
        backtrack(nums, track);
        // 取消选择
        track.removeLast();
    }
}
```

**Go**

```go
func permute(nums []int) [][]int {
    if len(result) != 0 {
        result = [][]int{}
    }
    backtrack(nums, []int{})
    return result
}

var result [][]int
// 路径：track，选择列表：nums中不在track中的部分
func backtrack(nums, track []int) {
    // 回溯终止条件
    if len(nums) == len(track) {
        // 特别注意：需要用一个新切片来保存结果，否则会在后续回溯中，破坏result里的结果
        // nTrack := make([]int, len(track))
        // copy(nTrack, track)
        // result = append(result, nTrack)

        // 或者简写成这样
        nTrack := append([]int{}, track...)
        result = append(result, nTrack)
        return
    }

    Flag:
    for i := 0; i < len(nums); i++ {
        // 过滤不合法选择
        for j := 0; j < len(track); j++ {
            if nums[i] == track[j] {
                continue Flag
            }
        }
        // 做选择 和 撤销选择 通过append同时实现
        track = append(track, nums[i])
        backtrack(nums, track)
        track = track[:len(track)-1]
    }
}
```

### 2. N皇后问题

#### 题目描述

n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

 ![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/8-queens.png) 

上图为 8 皇后问题的一种解法。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

**示例:**

```
输入: 4
输出: [
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法
```

#### 代码实现

**C++**

```c++
vector<vector<string>> res;

/* 输入棋盘边长 n，返回所有合法的放置 */
vector<vector<string>> solveNQueens(int n) {
    // '.' 表示空，'Q' 表示皇后，初始化空棋盘。
    vector<string> board(n, string(n, '.'));
    backtrack(board, 0);
    return res;
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
void backtrack(vector<string>& board, int row) {
    // 触发结束条件
    if (row == board.size()) {
        res.push_back(board);
        return;
    }

    int n = board[row].size();
    for (int col = 0; col < n; col++) {
        // 排除不合法选择
        if (!isValid(board, row, col)) 
            continue;
        // 做选择
        board[row][col] = 'Q';
        // 进入下一行决策
        backtrack(board, row + 1);
        // 撤销选择
        board[row][col] = '.';
    }
}
/* 是否可以在 board[row][col] 放置皇后？ */
bool isValid(vector<string>& board, int row, int col) {
    int n = board.size();
    // 检查列是否有皇后互相冲突
    for (int i = 0; i < n; i++) {
        if (board[i][col] == 'Q')
            return false;
    }
    // 检查右上方是否有皇后互相冲突
    for (int i = row - 1, j = col + 1; 
            i >= 0 && j < n; i--, j++) {
        if (board[i][j] == 'Q')
            return false;
    }
    // 检查左上方是否有皇后互相冲突
    for (int i = row - 1, j = col - 1;
            i >= 0 && j >= 0; i--, j--) {
        if (board[i][j] == 'Q')
            return false;
    }
    return true;
}
```

**Go**

```go
func solveNQueens(n int) [][]string {
    // LeetCode中，调用函数进行测试时，全局变量的值将一直保存，所以在每次调用函数时，需要手动清空
    if len(result) != 0 {
        result = [][]string{}
    }
    // 生成棋盘
    board := make([][]byte, n)
    b := []byte{}
    for i := 0; i < n; i++ {
        b = append(b, '.')
    }
    for i := 0; i < n; i++ {
        board[i] =  append([]byte{}, b...)
    }
    backtrack(board, 0)

    return result
}

var result [][]string
func backtrack(board [][]byte, row int) {
    // 终止条件
    if row == len(board) {
        b := bytes.Join(board, []byte(";"))
        strs := strings.Split(string(b), ";")
        result = append(result, strs)
        return
    }

    for col := 0; col < len(board); col++ {
        // 排除无效的选择
        if !isValid(board, row, col) { continue }
        // 做选择
        board[row][col] = 'Q'
        backtrack(board, row+1)
        // 撤销选择
        board[row][col] = '.'
    }
}

func isValid(board [][]byte, row, col int) bool {
    // 查找列
    // for i := 0; i < len(board); i++ {
    //    if i != row && board[i][col] == 'Q' {
    //        return false
    //    }
    // }
    
    // 查找列 截止到row的位置就行了
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }
    // 不需要查找整个 / 方向的元素，因为下半部分还没做出选择
    // 查找右上
    for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }
    // 同理，查找左上
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }

    return true
}
```

## 总结

**写`backtrack`函数时，需要维护走过的「路径」和当前可以做的「选择列表」，当触发「结束条件」时，将「路径」记入结果集**。

其实想想看，回溯算法和动态规划是不是有点像呢？我们在动态规划系列文章中多次强调，动态规划的三个需要明确的点就是「状态」「选择」和「base case」，是不是就对应着走过的「路径」，当前的「选择列表」和「结束条件」？

某种程度上说，动态规划的暴力求解阶段就是回溯算法。只是有的问题具有重叠子问题性质，可以用 dp table 或者备忘录优化，将递归树大幅剪枝，这就变成了动态规划。而今天的两个问题，都没有重叠子问题，也就是回溯算法问题了，复杂度非常高是不可避免的。

本文终。如果觉得本文不错的话，不妨分享给你的朋友。另外，几篇动态规划和框架思维的相关文章正在更新和修订，敬请期待。

> 原文始发于微信公众号（labuladong）：[回溯算法详解（修订版）](http://mp.weixin.qq.com/s/nMUHqvwzG2LmWA9jMIHwQQ)