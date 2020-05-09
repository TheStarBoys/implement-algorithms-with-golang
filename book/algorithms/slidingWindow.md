# 滑动窗口

## 简介

> 引用自：[面试官，你再问我滑动窗口问题试试？我有解题模板，不怕！](https://www.cxyxiaowu.com/672.html )
>

滑动窗口这类问题一般需要用到 **双指针** 来进行求解，另外一类比较特殊则是需要用到特定的数据结构，像是 sorted_map。

后者有特定的题型，后面会列出来，但是，对于前者，题形变化非常的大，一般都是基于字符串和数组的，所以我们重点总结这种基于双指针的滑动窗口问题。

题目问法大致有这几种：

- 给两个字符串，一长一短，问其中短的是否在长的中满足一定的条件存在，例如：
- 求长的的最短子串，该子串必须涵盖短的的所有字符
- 短的的 anagram 在长的中出现的所有位置
- …
- 给一个字符串或者数组，问这个字符串的子串或者子数组是否满足一定的条件，例如：
- 含有少于 k 个不同字符的最长子串
- 所有字符都只出现一次的最长子串
- …

除此之外，还有一些其他的问法，但是不变的是，这类题目脱离不开主串（主数组）和子串（子数组）的关系，要求的时间复杂度往往是 O(n) ，空间复杂度往往是常数级的。

之所以是滑动窗口，是因为，遍历的时候，两个指针一前一后夹着的子串（子数组）类似一个窗口，这个窗口大小和范围会随着前后指针的移动发生变化。

![面试官，你再问我滑动窗口问题试试？我有解题模板，不怕！](http://www.cxyxiaowu.com/wp-content/uploads/2019/10/1571057536-4a4dd87048cbf50.jpg)双指针确定一个窗口

## 解题思路及答题模板

根据前面的描述，滑动窗口就是这类题目的重点，换句话说，**窗口的移动** 就是重点！

我们要控制前后指针的移动来控制窗口，这样的移动是有条件的，也就是要想清楚在什么情况下移动，在什么情况下保持不变。

我的思路是保证右指针每次往前移动一格，每次移动都会有新的一个元素进入窗口，这时条件可能就会发生改变，然后根据当前条件来决定左指针是否移动，以及移动多少格。

我写来一个模版在这里，可以参考：

```java
//authour:P.yh
//Editor:程序员小吴
public int slidingWindowTemplate(String[] a, ...) {
    // 输入参数有效性判断
    if (...) {
        ...
    }

    // 申请一个散列，用于记录窗口中具体元素的个数情况
    // 这里用数组的形式呈现，也可以考虑其他数据结构
    int[] hash = new int[...];

    // 预处理(可省), 一般情况是改变 hash
    ...

    // l 表示左指针
    // count 记录当前的条件，具体根据题目要求来定义
    // result 用来存放结果
    int l = 0, count = ..., result = ...;
    for (int r = 0; r < A.length; ++r) {
        // 更新新元素在散列中的数量
        hash[A[r]]--;

        // 根据窗口的变更结果来改变条件值
        if (hash[A[r]] == ...) {
            count++;
        }

        // 如果当前条件不满足，移动左指针直至条件满足为止
        while (count > K || ...) {
            ...
            if (...) {
                count--;
            }
            hash[A[l]]++;
            l++;
        }

        // 更新结果
        results = ...
    }

    return results;
}
```

这里面的 “移动左指针直至条件满足” 部分，需要具体题目具体分析，其他部分的变化不大。

## 具体题目分析与代码

### 1. 找到字符串中所有字母异位词

题目来源于 LeetCode 上第 438 号问题：找到字符串中所有字母异位词。题目难度为 Easy，目前通过率为 43.6% 。

#### 题目描述

给定一个字符串 **s** 和一个非空字符串 **p**，找到 **s** 中所有是 **p** 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 **s** 和 **p** 的长度都不超过 **20100** 。

**说明：**

- 字母异位词指字母相同，但排列不同的字符串。
- 不考虑答案输出的顺序。

**示例 1:**

```
输入:
s: "cbaebabacd" p: "abc"

输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
```

 **示例 2:**

```
输入:
s: "abab" p: "ab"

输出:
[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
```

#### 题目解析

别看这是一道 easy 难度的题目，如果限定你在 O(n) 时间复杂度内实现呢？

按照模版会很简单！

首先窗口是固定的，窗口长度就是输入参数中第二个字符串的长度，也就是说，右指针移动到某个位置后，左指针必须跟着一同移动，且每次移动都是一格，模版中 count 用来记录窗口内满足条件的元素，直到 count 和窗口长度相等即可更新答案。

#### 代码实现

**Java**

```java
// 使用上面的模板进行解题，so easy ！
public List<Integer> findAnagrams(String s, String p) {
    // 输入参数有效性判断
    if (s.length() < p.length()) {
        return new ArrayList<Integer>();
    }

    // 申请一个散列，用于记录窗口中具体元素的个数情况
    // 这里用数组的形式呈现，也可以考虑其他数据结构
    char[] sArr = s.toCharArray();
    char[] pArr = p.toCharArray();

    int[] hash = new int[26];

    for (int i = 0; i < pArr.length; ++i) {
        hash[pArr[i] - 'a']++;
    }

    // l 表示左指针
    // count 记录当前的条件，具体根据题目要求来定义
    // result 用来存放结果
    List<Integer> results = new ArrayList<>();
    int l = 0, count = 0, pLength = p.length();

    for (int r = 0; r < sArr.length; ++r) {
        // 更新新元素在散列中的数量
        hash[sArr[r] - 'a']--;

        // 根据窗口的变更结果来改变条件值
        if (hash[sArr[r] - 'a'] >= 0) {
            count++;
        }

        // 如果当前条件不满足，移动左指针直至条件满足为止
        if (r > pLength - 1) {
            hash[sArr[l] - 'a']++;

            if (hash[sArr[l] - 'a'] > 0) {
                count--;
            }

            l++;
        }

        // 更新结果
        if (count == pLength) {
            results.add(l);
        }
    }

    return results;
}
```

**Go**

```go
func findAnagrams(s string, p string) []int {
    // 输入参数有效性判断
    if len(s) < len(p) {
        return []int{}
    }
    // 申请一个散列来记录窗口值
    hash := [26]int{}
    // l 左指针, count 记录和p中字符相同的字符个数, result 输出结果
    l, count, result := 0, 0, []int{}
    // 将p的值添加到窗口中
    for i := 0; i < len(p); i++ {
        hash[p[i] - 'a']++
    }
    // 开始移动窗口
    for r := 0; r < len(s); r++ {
        hash[s[r] - 'a']--
        // 如果窗口中拥有该元素
        if hash[s[r] - 'a'] >= 0 {
            count++
        }
        // 判断是否需要移动l指针
        if r - l > len(p) - 1 {
            hash[s[l] - 'a']++
            // 如果不是p中的字符，窗口中对应值最多会恢复到0
            // 如果是p中的字符，窗口对应值将大于0
            if hash[s[l] - 'a'] > 0 {
                count--
            }
            l++
        }
        // 更新结果
        if count == len(p) {
            result = append(result, l)
        }
    }

    return result
}
```

### 2. 最小覆盖子串

题目来源于 LeetCode 上第 76 号问题：最小覆盖子串。题目难度为 Hard，目前通过率为 35.8% 。

#### 题目描述

给你一个字符串 **S**、一个字符串 **T**，请在字符串 **S** 里面找出：包含 **T** 所有字母的最小子串。

**示例：**

```java
输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
```

#### 题目解析

同样是两个字符串之间的关系问题，因为题目求的最小子串，也就是窗口的最小长度，说明这里的窗口大小是可变的，这里移动左指针的条件变成，只要左指针指向不需要的字符，就进行移动。

依旧使用上面的模板解题！

#### 代码实现 

**Java**

```java
// 修改了原著中的变量名max ==> min
// 使用上面的模板进行解题，受篇幅限制下面的代码就不添加注释了
public String minWindow(String s, String t) {
    if (s.length() < t.length()) {
        return "";
    }

    char[] sArr = s.toCharArray();
    char[] tArr = t.toCharArray();

    int[] hash = new int[256];

    for (int i = 0; i < tArr.length; ++i) {
        hash[tArr[i]]++;
    }

    int l = 0, count = tArr.length, min = s.length() + 1;
    String result = "";
    for (int r = 0; r < sArr.length; ++r) {
        hash[sArr[r]]--;

        if (hash[sArr[r]] >= 0) {
            count--;
        }

        while (l < r && hash[sArr[l]] < 0) {
            hash[sArr[l]]++;
            l++;
        }

        if (count == 0 && min > r - l + 1) {
            min = r - l + 1;
            result = s.substring(l, r + 1);
        }
    }

    return result;
}
```

**Go**

```go
func minWindow(s string, t string) string {
    // 输出有效性判断
    if len(s) < len(t) {
        return ""
    }
    hash := [256]int{}
    for i := 0; i < len(t); i++ {
        hash[t[i]]++
    }

    // 如果包含T所有字符的最小子串刚好就是S的话，那么min的边界条件就是需要大于S的长度
    l, count, result, min := 0, len(t), "", len(s) + 1
    for r := 0; r < len(s); r++ {
        hash[s[r]]--
        if hash[s[r]] >= 0 {
            count--
        }
        // 如果是多余或不需要的元素，将移动左指针
        for l < r && hash[s[l]] < 0 {
            hash[s[l]]++
            l++
        }
        if count == 0 && min > r - l + 1 {
            min = r - l + 1
            result = s[l:r+1]
        }
    }

    return result
}
```



###  3. 无重复字符的最长子串 

题目来源于 LeetCode 上第 3 号问题：无重复字符的最长子串。题目难度为 Medium，目前通过率为 29.0% 。

#### 题目描述

给定一个字符串，请你找出其中不含有重复字符的 **最长子串** 的长度。

**示例 1:**

```java
输入: "abcabcbb"输出: 3 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

```

#### 题目解析

输入只有一个字符串，要求子串里面不能够有重复的元素，这里 count 都不需要定义，直接判断哈希散列里面的元素是不是在窗口内即可，是的话得移动左指针去重。

具体操作如下：

建立一个 256 位大小的整型数组 freg ，用来建立字符和其出现位置之间的映射。

维护一个滑动窗口，窗口内的都是没有重复的字符，去尽可能的扩大窗口的大小，窗口不停的向右滑动。

- （1）如果当前遍历到的字符从未出现过，那么直接扩大右边界；
- （2）如果当前遍历到的字符出现过，则缩小窗口（左边索引向右移动），然后继续观察当前遍历到的字符；
- （3）重复（1）（2），直到左边索引无法再移动；
- （4）维护一个结果 res，每次用出现过的窗口大小来更新结果 res，最后返回 res 获取结果。

#### 动画描述

![面试官，你再问我滑动窗口问题试试？我有解题模板，不怕！](http://www.cxyxiaowu.com/wp-content/uploads/2019/10/1571057537-818020d836023ea.gif)

#### 代码实现

**Java**

```java
public int lengthOfLongestSubstring(String s) {
    if (s == null || s.length() == 0) {
        return 0;
    }

    char[] sArr = s.toCharArray();
    int[] hash = new int[256];

    int l = 0, result = 1;
    for (int r = 0; r < sArr.length; ++r) {
        hash[sArr[r]]++;

        while (hash[sArr[r]] != 1) {
            hash[sArr[l]]--;
            l++;
        }

        result = Math.max(result, r - l + 1);
    }

    return result;
}
```

**Go**

```go
func lengthOfLongestSubstring(s string) int {
    hash := [256]int{}
    l, max := 0, 0

    for r := 0; r < len(s); r++ {
        hash[s[r]]++
        // 出现了重复元素
        for hash[s[r]] > 1 && l < r {
            hash[s[l]]--
            l++
        }
        if max < r - l + 1 {
            max = r - l + 1
        }
    }

    return max
}
```



### 4. 字符串的排列

题目来源于 LeetCode 上第 567 号问题：字符串的排列。题目难度为 Medium，目前通过率为 31.8% 。

#### 题目描述

给定两个字符串 **s1** 和 **s2**，写一个函数来判断 **s2** 是否包含 **s1** 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

**示例1:**

```
输入: s1 = "ab" s2 = "eidbaooo"输出: True解释: s2 包含 s1 的排列之一 ("ba").
```

**示例2:**

```
输入: s1= "ab" s2 = "eidboaoo"输出: False
```

#### 题目解析

和 438 那题很类似，但是这里不需要记录答案了，有就直接返回 true。

#### 代码实现

**Java**

```java
public boolean checkInclusion(String s1, String s2) {
    if (s1.length() > s2.length()) {
        return false;
    }

    char[] s1Arr = s1.toCharArray();
    char[] s2Arr = s2.toCharArray();

    int[] hash = new int[26];

    for (int i = 0; i < s1Arr.length; ++i) {
        hash[s1Arr[i] - 'a']++;
    }

    int l = 0, count = 0;
    for (int r = 0; r < s2Arr.length; ++r) {
        hash[s2Arr[r] - 'a']--;

        if (hash[s2Arr[r] - 'a'] >= 0) {
            count++;
        }

        if (r >= s1Arr.length) {
            hash[s2Arr[l] - 'a']++;

            if (hash[s2Arr[l] - 'a'] >= 1) {
                count--;
            }

            l++;
        }

        if (count == s1Arr.length) {
            return true;
        }
    }

    return false;
}
```

**Go**

```go
func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    hash := [256]int{}
    l, count := 0, 0
    // 导入s1
    for i := 0; i < len(s1); i++ {
        hash[s1[i]]++
    }
    for r := 0; r < len(s2); r++ {
        hash[s2[r]]--
        // 根据hash改变count条件值
        if hash[s2[r]] >= 0 {
            count++
        } else {
            // 更新l
            for hash[s2[r]] < 0 {
                hash[s2[l]]++
                // 根据hash改变count条件值
                if hash[s2[l]] > 0 {
                    count--
                }
                l++
            }
        }

        if count == len(s1) {
            return true
        }
    }

    return false
}
```



### 5. K 个不同整数的子数组

题目来源于 LeetCode 上第 992 号问题： K 个不同整数的子数组。题目难度为 Hard，目前通过率为 26.4% 。

#### 题目描述

给定一个正整数数组 **A**，如果 **A** 的某个子数组中不同整数的个数恰好为 **K**，则称 **A** 的这个连续、不一定独立的子数组为好子数组。

（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）

返回 **A **中好子数组的数目。

**示例 1：**

```
输出：A = [1,2,1,2,3], K = 2输入：7解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
```

**示例 2：**

```
输入：A = [1,2,1,3,4], K = 3输出：3解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].
```

#### 题目解析

看完了字符串类型的题目，这次来看看数组类型的。

题目中的 subarray 已经明确了这个题可以考虑用滑动窗口，这题比较 trick 的一个地方在于，这里不是求最小值最大值，而是要你计数。

但是如果每次仅仅加 1 的话又不太对，例如 `A = [1,2,1,2,3], K = 2` 这个例子，假如右指针移到 index 为 3 的位置，如果按之前的思路左指针根据 count 来移动，当前窗口是 `[1,2,1,2]`，但是怎么把 `[2,1]` 给考虑进去呢？

可以从数组和子数组的关系来思考！

假如 `[1,2,1,2]` 是符合条件的数组，如果要计数的话，`[1,2,1,2]` 要求的结果是否和 `[1,2,1]` 的结果存在联系？这两个数组的区别在于多了一个新进来的元素，之前子数组计数没考虑到这个元素，假如把这个元素放到之前符合条件的子数组中组成的新数组也是符合条件的，我们看看这个例子中所有满足条件的窗口以及对应的满足条件的子数组情况：

```java
[1,2,1,2,3]  // 窗口满足条件
 l r         // 满足条件的子数组 [1,2]

[1,2,1,2,3]  // 窗口满足条件
 l   r       // 满足条件的子数组 [1,2],[2,1],[1,2,1]

[1,2,1,2,3]  // 窗口满足条件
 l     r     // 满足条件的子数组 [1,2],[2,1],[1,2,1],[1,2],[2,1,2],[1,2,1,2]

[1,2,1,2,3]  // 窗口不满足条件，移动左指针至满足条件
 l       r   

[1,2,1,2,3]  // 窗口满足条件
       l r   // 满足条件的子数组 [1,2],[2,1],[1,2,1],[1,2],[2,1,2],[1,2,1,2],[2,3]
```

你可以看到对于一段连续的数组，新的元素进来，窗口增加 1，每次的增量都会在前一次增量的基础上加 1。

当新的元素进来打破当前条件会使这个增量从新回到 1，这样我们左指针移动条件就是只要是移动不会改变条件，就移动，不然就停止。

#### 代码实现

```java
public int subarraysWithKDistinct(int[] A, int K) {
    if (A == null || A.length < K) {
        return 0;
    }

    int[] hash = new int[A.length + 1];

    int l = 0, results = 0, count = 0, result = 1;
    for (int r = 0; r < A.length; ++r) {
        hash[A[r]]++;

        if (hash[A[r]] == 1) {
            count++;
        }

        while (hash[A[l]] > 1 || count > K) {
            if (count > K) {
                result = 1;
                count--;
            } else {
                result++;
            }
            hash[A[l]]--;
            l++;
        }

        if (count == K) {
            results += result;
        }
    }

    return results;
}
```

**Go**

```go
// first 超时，需要优化
func subarraysWithKDistinct(A []int, K int) int {
    if len(A) < K {
        return 0
    }
    hash := make(map[int]int)
    var result int
    for i := 0; i < len(A); i++ {
        l := i
        // 注意r 是从l处开始的
        for r := l; r < len(A); r++ {
            hash[A[r]]++
            if hash[A[r]] == 1 {
                K--
            }
            if K == 0 {
                // fmt.Println(A[l:r+1])
                result++
            }
        }

        for r := len(A) - 1; r >= l; r-- {
            hash[A[r]]--
            if hash[A[r]] == 0 {
                K++
            }
        }
    }
    return result
}

// second 作者思路，依旧超时
func subarraysWithKDistinct(A []int, K int) int {
    if len(A) < K {
        return 0
    }
    // 题目给了正整数的取值范围
    // 数组声明不支持用非常量定义数组大小
    // hash := [len(A)+1]int{}
    hash := make([]int, len(A)+1, len(A)+1)
    // results 为最终结果, result 为前一次的增量，count为不重复的整数个数
    l, results, count, result := 0, 0, 0, 1
    for r := 0; r < len(A); r++ {
        hash[A[r]]++
        // 统计不重复的整数个数
        if hash[A[r]] == 1 {
            count++
        }
        for hash[A[l]] > 1 || count > K {
            fmt.Printf("before, A: %v, result: %d\n", A[l:r+1], result)
            if count > K {
                count--
                result = 1
            } else {
                result++
            }
            hash[A[l]]--
            l++
            fmt.Printf("after, A: %v, result: %d\n", A[l:r+1], result)
        }
        if count == K {
            results += result
            fmt.Printf("A:%v, results: %d\n", A[l:r+1], results)
        }
    }
    
    return results
}
```



### 6. 替换后的最长重复字符

题目来源于 LeetCode 上第 424 号问题：替换后的最长重复字符。题目难度为 Medium，目前通过率为 37.3% 。

#### 题目描述

给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 *k* 次。在执行上述操作后，找到包含重复字母的最长子串的长度。

#### 题目解析

这道题想 accept 的话不难，但是问题在于怎么知道当前窗口中数量最多的字符的数量，因为需要替换的字符就是当前窗口的大小减去窗口中数量最多的字符的数量。

最简单的方法就是把 **哈希散列** 遍历一边找到最大的字符数量，但是仔细想想如果我们每次新进元素都更新这个最大数量，且只更新一次，我们保存的是当前遍历过的全局的最大值，它肯定是比实际的最大值大的，我们左指针移动的条件是 `r - l + 1 - maxCount > k`，保存的结果是 `result = Math.max(r - l + 1, result);` 这里 `maxCount` 比实际偏大的话，虽然导致左指针不能移动，但是不会记录当前的结果，所以最后的答案并不会受影响。

#### 代码实现

**Java**

```java
public int characterReplacement(String s, int k) {
    if (s == null || s.length() == 0) {
        return 0;
    }

    char[] sArr = s.toCharArray();

    int[] hash = new int[26];

    int l = 0, maxCount = 0, result = 0;
    for (int r = 0; r < sArr.length; ++r) {
        hash[sArr[r] - 'A']++;

        maxCount = Math.max(maxCount, hash[sArr[r] - 'A']);

        while (r - l + 1 - maxCount > k) {
            hash[sArr[l] - 'A']--;
            l++;
        }

        result = Math.max(r - l + 1, result);
    }

    return result;
}
```

**Go**

```go
func characterReplacement(s string, k int) int {
    // 转换题意：
    // 在窗口中，最多有K个不同的字符，求窗口最大长度
    // 在窗口中，最少有W - K个相同的字符，求窗口最大长度, W为窗口大小
    if k >= len(s) {
        return len(s)
    }
    hash := [26]int{}
    // count为最大相同字符数，max为结果
    l, count, max := 0, 1, 0
    for r := 0; r < len(s); r++ {
        hash[s[r] - 'A']++
        // 更新最大字符数
        if hash[s[r] - 'A'] > count {
            count = hash[s[r] - 'A']
        }
        // l < r可以省略，因为count最小值为0，那么则有 r - l + 1 > k 成立时，循环才会执行
        // k最小为0，则有 r + 1 > l时才会执行，所有 l 是肯定小于 r 的
        // for count < (r - l + 1) - k && l < r {
        for count < (r - l + 1) - k {
            hash[s[l] - 'A']--
            l++
        }
        // max 只跟窗口大小有关，只要 l 跟 r 值正确，这一步就正确
        max = int(math.Max(float64(max), float64(r - l + 1)))
        
    }

    return max
}
```



## 总结

双指针类的滑动窗口问题思维复杂度并不高，但是出错点往往在细节。

记忆常用的解题模版还是很有必要的，特别是对于这种变量名多，容易混淆的题型。

有了这个框架，思考的点就转化为 “什么条件下移动左指针”，无关信息少了，思考加实现自然不是问题。