# 数据结构与算法

## 简介

> 文中引用均已说明出处，侵删。

- 主要用Go实现的数据结构与算法，文中引用的文章都非常不错，建议作为扩展阅读。

- 本项目是收集各类文章所作出的一个简单的笔记，方便复习。
- 为苦于寻找Go相应的数据结构与算法实现的Gopher提供一个参考的实现。
- 如果项目中有疏漏，欢迎指正。



## 思想

对于一些算法，平时工作几乎不可能遇到，那为何还要学？

- 优化意识
  - 算法诞生的原因？好处和坏处是什么？
  - 查找时减少时间复杂度——散列表
  - 某表达式开销大——预处理，并缓存
  - 负载均衡——通过哈希进行分片
- 思维训练
  - 写代码时不由自主考虑性能问题
  - 降低代码出现bug的几率
  - 特定的算法依赖于特定的数据结构
  - 练拳不练功到老一场空
- 应用
  - 学习它的由来、特性、适用场景以及它能解决的问题

## 目录

### 入门

- 复杂度分析

### 数据结构

- [栈与队列](./book/dataStructure/stackAndQueue.md)
- [链表](./book/dataStructure/linkedList.md)
- [二叉树](./book/dataStructure/binaryTree.md)
- [二叉搜索树](./book/dataStructure/binarySearchTree.md)
- [N叉树](./book/dataStructure/n-aryTree.md)
- [前缀树](./book/dataStructure/trie.md)
- [散列表](./book/dataStructure/hashTable.md)
- 跳表
- [堆](./book/dataStructure/heap.md)
- [图](https://time.geekbang.org/column/article/70537)

### 算法

- [双指针](./book/algorithms/doublePointer.md)
- [递归](./book/algorithms/recursiveAlgorithm.md)
- [排序](./book/algorithms/sort.md)
- [二分查找](./book/algorithms/binarySearch.md)
- [贪心算法](./book/algorithms/greedyAlgorithm.md)
- [动态规划](./book/algorithms/dynamicProgramming.md)
- [滑动窗口](./book/algorithms/slidingWindow.md)
- [回溯算法](./book/algorithms/backtrack.md)
- [分治算法](./book/algorithms/divideAndConquerAlgorithm.md)
- [广度优先搜索](./book/algorithms/bfs.md)
- [深度优先搜索](./book/algorithms/dfs.md)
- [字符串匹配](./book/algorithms/stringMatching.md)

### Go源码分析

- [container/list](./book/dataStructure/container_list.md)

