# 前缀树

## 简介

> 引用自：[LeetCode](https://leetcode-cn.com/explore/learn/card/trie/)
>
> 代码实现：TheStarBoys

`前缀树`，又称`字典树`，是`N叉树`的特殊形式。

在这章中，我们将深入讨论前缀树的实现方法以及如何将这个数据结构应用到实际问题中。

完成这章后，你将：

1. 理解前缀树的`基本概念`；
2. 掌握前缀树中的`插入`和`搜索操作`；
3. 了解前缀树如何帮助解决`实际应用问题`；
4. 运用前缀树解题。

此章至少需要的预备知识：

- [N叉树](./n-aryTree.md)

## 概念

`前缀树`是`N叉树`的一种特殊形式。通常来说，一个前缀树是用来`存储字符串`的。前缀树的每一个节点代表一个`字符串`（`前缀`）。每一个节点会有多个子节点，通往不同子节点的路径上有着不同的字符。子节点代表的字符串是由节点本身的`原始字符串`，以及`通往该子节点路径上所有的字符`组成的。

下面是前缀树的一个例子：

![img](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/02/07/screen-shot-2018-01-31-at-163403.png)

在上图示例中，我们在节点中标记的值是该节点对应表示的字符串。例如，我们从根节点开始，选择第二条路径 'b'，然后选择它的第一个子节点 'a'，接下来继续选择子节点 'd'，我们最终会到达叶节点 "bad"。节点的值是由从根节点开始，与其经过的路径中的字符按顺序形成的。

值得注意的是，根节点表示`空字符串`。

前缀树的一个重要的特性是，节点所有的后代都与该节点相关的字符串有着共同的前缀。这就是`前缀树`名称的由来。

我们再来看这个例子。例如，以节点 "b" 为根的子树中的节点表示的字符串，都具有共同的前缀 "b"。反之亦然，具有公共前缀 "b" 的字符串，全部位于以 "b" 为根的子树中，并且具有不同前缀的字符串来自不同的分支。

前缀树有着广泛的应用，例如自动补全，拼写检查等等。我们将在后面的章节中介绍实际应用场景。

## 表示方法

 前缀树的特别之处在于字符和子节点之间的对应关系。有许多不同的表示前缀树节点的方法，这里我们只介绍其中的两种方法。 

### 方法一：数组

------

第一种方法是用`数组`存储子节点。

例如，如果我们只存储含有字母 `a` 到 `z` 的字符串，我们可以在每个节点中声明一个大小为26的数组来存储其子节点。对于特定字符 `c`，我们可以使用 `c - 'a'` 作为索引来查找数组中相应的子节点。

```go
// change this value to adapt to different cases
const N = 26
type TrieNode {
    children [N]*TrieNode
    // you might need some extra values according to different cases
}
```

 访问子节点十分`快捷`。访问一个特定的子节点比较`容易`，因为在大多数情况下，我们很容易将一个字符转换为索引。但并非所有的子节点都需要这样的操作，所以这可能会导致`空间的浪费`。 

### 方法二 - Map

------

第二种方法是使用 `Hashmap` 来存储子节点。

我们可以在每个节点中声明一个Hashmap。Hashmap的键是字符，值是相对应的子节点。

```go
type TrieNode {
    children map[byte]*TrieNode
    // you might need some extra values according to different cases
}
```

 通过相应的字符来访问特定的子节点`更为容易`。但它可能比使用数组`稍慢一些`。但是，由于我们只存储我们需要的子节点，因此`节省了空间`。这个方法也更加`灵活`，因为我们不受到固定长度和固定范围的限制。 

### 补充

------

我们已经提到过如何表示前缀树中的子节点。除此之外，我们也需要用到一些其他的值。

例如，我们知道，前缀树的每个节点表示一个字符串，但并不是所有由前缀树表示的字符串都是有意义的。如果我们只想在前缀树中存储单词，那么我们可能需要在每个节点中声明一个布尔值（Boolean）作为标志，来表明该节点所表示的字符串是否为一个单词。



## 基本操作

### 在前缀树中插入

我们已经在另一章中讨论了 ([如何在二叉搜索树中实现插入操作](./binaryTree.md))。

> 提问：
>
> 你还记得如何在二叉搜索树中插入一个新的节点吗？

当我们在二叉搜索树中插入目标值时，在每个节点中，我们都需要根据 `节点值` 和 `目标值` 之间的关系，来确定目标值需要去往哪个子节点。同样地，当我们向前缀树中插入一个目标值时，我们也需要根据插入的 `目标值` 来决定我们的路径。

更具体地说，如果我们在前缀树中插入一个字符串 `S`，我们要从根节点开始。 我们将根据 `S[0]`（S中的第一个字符），选择一个子节点或添加一个新的子节点。然后到达第二个节点，并根据 `S[1]` 做出选择。 再到第三个节点，以此类推。 最后，我们依次遍历 S 中的所有字符并到达末尾。 末端节点将是表示字符串 S 的节点。



我们来用伪代码总结一下以上策略：

```
1. Initialize: cur = root
2. for each char c in target string S:
3.      if cur does not have a child c:
4.          cur.children[c] = new Trie node
5.      cur = cur.children[c]
6. cur is the node which represents the string S
```

 通常情况情况下，你需要自己构建前缀树。构建前缀树实际上就是多次调用插入函数。但请记住在插入字符串之前要 `初始化根节点` 。 

### 在前缀树中搜索

#### 搜索前缀

正如我们在前缀树的简介中提到的，所有节点的后代都与该节点相对应字符串的有着共同前缀。因此，很容易搜索以特定前缀开头的任何单词。

同样地，我们可以根据给定的前缀沿着树形结构搜索下去。一旦我们找不到我们想要的子节点，搜索就以失败终止。否则，搜索成功。为了更具体地解释搜索的过程，我们提供了下列示例：（如果需要看动图，可以到LeetCode去看）

我们来用伪代码总结一下以上策略：

```
1. Initialize: cur = root
2. for each char c in target string S:
3.      if cur does not have a child c:
4.          search fails
5.      cur = cur.children[c]
6. search successes
```

#### 搜索单词

你可能还想知道如何搜索特定的单词，而不是前缀。我们可以将这个词作为前缀，并同样按照上述同样的方法进行搜索。

1. 如果搜索失败，那么意味着没有单词以目标单词开头，那么目标单词绝对不会存在于前缀树中。
2. 如果搜索成功，我们需要检查目标单词是否是前缀树中单词的前缀，或者它本身就是一个单词。为了进一步解决这个问题，你可能需要稍对节点的结构做出修改。

> 提示：往每个节点中加入布尔值可能会有效地帮助你解决这个问题。



### 实现Trie（前缀树）

#### 题目描述

实现一个 Trie (前缀树)，包含 `insert`, `search`, 和 `startsWith` 这三个操作。

**示例:**

```
Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");   
trie.search("app");     // 返回 true
```

**说明:**

- 你可以假设所有的输入都是由小写字母 `a-z` 构成的。
- 保证所有输入均为非空字符串。

#### 代码实现

```go
type Trie struct {
    children map[byte]*Trie
    isWord bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        children: make(map[byte]*Trie),
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    trie := this
    for i := range word {
        next, ok := trie.children[word[i]]
        if !ok {
            tmp := Constructor()
            next = &tmp
            trie.children[word[i]] = next
        }
        trie = next
    }
    trie.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    trie := this
    for i := range word {
        next, ok := trie.children[word[i]]
        if !ok {
            return false
        }
        trie = next
    }
    
    return trie.isWord || len(trie.children) == 0
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    trie := this
    for i := range prefix {
        next, ok := trie.children[prefix[i]]
        if !ok {
            return false
        }
        trie = next
    }
    
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
```

## 实际应用Ⅰ

### Map Sum Pairs

#### 题目描述

实现一个 MapSum 类里的两个方法，`insert` 和 `sum`。

对于方法 `insert`，你将得到一对（字符串，整数）的键值对。字符串表示键，整数表示值。如果键已经存在，那么原来的键值对将被替代成新的键值对。

对于方法 `sum`，你将得到一个表示前缀的字符串，你需要返回所有以该前缀开头的键的值的总和。

**示例 1:**

```
输入: insert("apple", 3), 输出: Null
输入: sum("ap"), 输出: 3
输入: insert("app", 2), 输出: Null
输入: sum("ap"), 输出: 5
```

#### 代码实现

```go
type MapSum struct {
    children map[byte]*MapSum
    val int
}


/** Initialize your data structure here. */
func Constructor() MapSum {
    return MapSum{
        children: make(map[byte]*MapSum),
    }
}


func (this *MapSum) Insert(key string, val int)  {
    trie := this
    for i := range key {
        next, ok := trie.children[key[i]]
        if !ok {
            tmp := Constructor()
            next = &tmp
            trie.children[key[i]] = next
        }
        trie = next
    }
    trie.val = val
}


func (this *MapSum) Sum(prefix string) int {
    // find prefix trie
    trie := this
    for i := range prefix {
        next, ok := trie.children[prefix[i]]
        if !ok {
            return 0
        }
        trie = next
    }
    sum := 0
    queue := make([]*MapSum, 0)
    queue = append(queue, trie)
    for len(queue) > 0 {
        length := len(queue)
        for i := 0; i < length; i++ {
            cur := queue[0]
            queue = queue[1:]
            sum += cur.val
            for _, node := range cur.children {
                queue = append(queue, node)
            }
        }
    }
    
    return sum
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
```

### 单词替换

#### 题目描述

在英语中，我们有一个叫做 `词根`(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 `继承词`(successor)。例如，词根`an`，跟随着单词 `other`(其他)，可以形成新的单词 `another`(另一个)。

现在，给定一个由许多词根组成的词典和一个句子。你需要将句子中的所有`继承词`用`词根`替换掉。如果`继承词`有许多可以形成它的`词根`，则用最短的词根替换它。

你需要输出替换之后的句子。

 

**示例：**

```
输入：dict(词典) = ["cat", "bat", "rat"] sentence(句子) = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
```

 

**提示：**

- 输入只包含小写字母。
- `1 <= dict.length <= 1000`
- `1 <= dict[i].length <= 100`
- 1 <= 句中词语数 <= 1000
- 1 <= 句中词语长度 <= 1000

#### 代码实现

```go
func replaceWords(dict []string, sentence string) string {
    trie := Constructor()
    for _, d := range dict {
        trie.Insert(d)
    }
    strs := strings.Split(sentence, " ")
    for i, word := range strs {
        root := trie.Judge(word)
        if root != word {
            strs[i] = root
        }
    }
    return strings.Join(strs, " ")
}

type Trie struct {
    children map[byte]*Trie
    isRoot bool
}

func Constructor() *Trie {
    return &Trie{
        children: make(map[byte]*Trie),
    }
}

func (this *Trie) Insert(root string) {
    trie := this
    for i := range root {
        next, ok := trie.children[root[i]]
        if !ok {
            next = Constructor()
            trie.children[root[i]] = next
        }
        trie = next
    }
    trie.isRoot = true
}

// 如果是继承词，返回其词根；如果不是继承词，返回其本身
func (this *Trie) Judge(word string) string {
    trie := this
    for i := range word {
        next, ok := trie.children[word[i]]
        if !ok {
            return word
        }
        if next.isRoot {
            return word[:i+1]
        }
        trie = next
    }
    return word
}
```



### 添加与搜索单词 - 数据结构设计

#### 题目描述

设计一个支持以下两种操作的数据结构：

```
void addWord(word)
bool search(word)
```

search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 `.` 或 `a-z` 。 `.` 可以表示任何一个字母。

**示例:**

```
addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
```

**说明:**

你可以假设所有单词都是由小写字母 `a-z` 组成的。

#### 代码实现

**方法一：map作为孩子**

```go
type WordDictionary struct {
    children map[byte]*WordDictionary
    isWord bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{
        children: make(map[byte]*WordDictionary),
    }
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
    trie := this
    for i := range word {
        next, ok := trie.children[word[i]]
        if !ok {
            tmp := Constructor()
            next = &tmp
            trie.children[word[i]] = next
        }
        trie = next
    }
    trie.isWord = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    trie := this
    for i := range word {
        // 对于 '.' , 递归的判断所有不为空的孩子
        if word[i] == '.' {
            for _, node := range trie.children {
                if node.Search(word[i+1:]) {
                    return true
                }
            }
            return false
        }
        next, ok := trie.children[word[i]]
        if !ok {
            return false
        }
        trie = next
    }
    return trie.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
```

**方法二：数组作为孩子**

```go
type WordDictionary struct {
    children [26]*WordDictionary
    isWord bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
    trie := this
    for i := range word {
        next := trie.children[word[i] - 'a']
        if next == nil {
            tmp := Constructor()
            next = &tmp
            trie.children[word[i] - 'a'] = next
        }
        trie = next
    }
    trie.isWord = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
    trie := this
    for i := range word {
        // 对于 '.' , 递归的判断所有不为空的孩子
        if word[i] == '.' {
            for _, node := range trie.children {
                if node == nil {
                    continue
                }
                if node.Search(word[i+1:]) {
                    return true
                }
            }
            return false
        }
        next := trie.children[word[i] - 'a']
        if next == nil {
            return false
        }
        trie = next
    }
    return trie.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
```



## 实际应用Ⅱ

