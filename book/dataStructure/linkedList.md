# 链表

## 简介
```go
type ListNode struct {
	next  *ListNode
	value interface{}
}
```

## 练习

### 1. [两数相加](https://leetcode-cn.com/problems/add-two-numbers)
#### 题目描述
给你两个 **非空** 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例1:
```
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
```
示例2:
```
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
```

#### 答案
```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        return nil
    }

    head := &ListNode{}
    curr := head
    mark := 0
    for l1 != nil || l2 != nil || mark == 1 {
        v1, v2 := 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }

        res := v1 + v2 + mark
        if res > 9 {
            mark = 1
        } else {
            mark = 0
        }
        curr.Val = res % 10
        if l1 != nil || l2 != nil || mark == 1 {
            curr.Next = &ListNode{}
            curr = curr.Next
        }
    }

    return head
}
```

### 2. [删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list)
#### 题目描述
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
示例1:
```
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
```
示例2:
```
输入：head = [1], n = 1
输出：[]
```
示例3:
```
输入：head = [1,2], n = 1
输出：[1]
```
#### 答案
解法1：两次扫描，第一次获取链表长度

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
    for tmp := head; tmp != nil; tmp = tmp.Next {
        length++
    }
    // 删除元素是正数第一个
    if n == length {
        return head.Next
    }
    // 找到要删除节点的前一个节点 tmp
    i, tmp := 0, head
    for ; i < length - n - 1; i++ {
        tmp = tmp.Next
    }
    tmp.Next = tmp.Next.Next

    return head
}
```
解法2：一次扫描，空间换时间
```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    arr := []*ListNode{}
    for ; head != nil; head = head.Next {
        arr = append(arr, head)
    }
    if len(arr) == n { return arr[0].Next }
    prevNode := arr[len(arr)-n-1]
    prevNode.Next = prevNode.Next.Next
    return arr[0]
}
```

### 3. [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists)
#### 题目描述
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例1:
```
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
```
示例2:
```
输入：l1 = [], l2 = []
输出：[]
```
示例3:
```
输入：l1 = [], l2 = [0]
输出：[0]
```
**提示**
- 两个链表的节点数目范围是 [0, 50]
- -100 <= Node.val <= 100
- l1 和 l2 均按 非递减顺序 排列
#### 答案
```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
        return nil
    }

    head := &ListNode{}
    curr := head
    for list1 != nil || list2 != nil {
        v1, v2 := math.MaxInt64, math.MaxInt64
        if list1 != nil {
            v1 = list1.Val
        }

        if list2 != nil {
            v2 = list2.Val
        }

        if v1 <= v2 {
            curr.Val = v1
            list1 = list1.Next
        } else {
            curr.Val = v2
            list2 = list2.Next
        }

        if list1 != nil || list2 != nil {
            curr.Next = &ListNode{}
            curr = curr.Next
        }
    }

    return head
}
```

### 4. [反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

#### 题目描述

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)

```
输入：head = [1,2]
输出：[2,1]
```

**示例 3：**

```
输入：head = []
输出：[]
```

**提示**

- 链表中节点的数目范围是 `[0, 5000]`
- `-5000 <= Node.val <= 5000`

**进阶：**链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

#### 答案

解法1：迭代

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }

    return prev
}
```



解法2：递归

```go
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}
```



### 5. [回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/)

#### 题目描述

给你一个单链表的头节点 `head` ，请你判断该链表是否为回文链表。如果是，返回 `true` ；否则，返回 `false` 。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/03/03/pal1linked-list.jpg)

```
输入：head = [1,2,2,1]
输出：true
```

示例 2：

![img](https://assets.leetcode.com/uploads/2021/03/03/pal2linked-list.jpg)

输入：head = [1,2]
输出：false

**提示**：

- 链表中节点数目在范围`[1, 10^5]` 内

- `0 <= Node.val <= 9`

#### 答案

解法1：链表转数组

```go
func isPalindrome(head *ListNode) bool {
    arr := []*ListNode{}
    for ; head != nil; head = head.Next {
        arr = append(arr, head)
    }

    for l, r := 0, len(arr) - 1; l < r; l, r = l+1, r-1 {
        if arr[l].Val != arr[r].Val {
            return false
        }
    }

    return true
}
```

解法2：通过双指针找到中间位置，将中间位置前半部分的链表反转，再遍历比较

```go
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil { return true }
    slow, fast := head.Next, head.Next.Next
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    var pre *ListNode
    var next *ListNode
    for head != slow {
        next = head.Next
        head.Next = pre
        pre = head
        head = next
    }
    if fast != nil { slow = slow.Next }
    for pre != nil {
        if pre.Val != slow.Val { return false }
        pre = pre.Next
        slow = slow.Next
    }
    return true
}
```

