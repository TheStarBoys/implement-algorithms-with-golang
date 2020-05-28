# 二分查找

## 简介

二分查找其实就是选择一个区间的中间值，首先判断该中间值是否是目标值，如果不满足，再根据一定的条件，判断接下来是往中间值的左侧去找目标值，还是往右侧去找目标值。这样的操作，将查找的区间每次都分为了三个部分。左侧，中间值，右侧。极大的提高了查找的效率。

## 模版Ⅰ

```go
func binarySearch(nums []int) int {
    // 根据题目可删减
    if len(nums) == 0 {
        return -1
    }
    left, right := 0, len(nums) - 1
    for left <= right {
		// 一般写法：mid := (left + right) / 2
		// 推荐写法：mid := left + (right - left) / 2,避免(left + right)溢出
		// 极致优化：mid := left + ((right - left) >> 1) 相比除法运算来说，计算机处理位运算要快得多
		mid := left + ((right - left) >> 1)
		if mid位置就是需要查找的数 {
			return mid
		} else if 如果待查找数在右侧 {
            left = mid + 1
        } else {
            // 如果待查找数在左侧
            right = mid - 1
        }
    }
    
    // 没有找到该数
    return -1
}
```

模板 #1 是二分查找的最基础和最基本的形式。这是一个标准的二分查找模板，大多数高中或大学会在他们第一次教学生计算机科学时使用。模板 #1 用于查找可以通过*访问数组中的单个索引*来确定的元素或条件。 

**关键属性**

- 二分查找的最基础和最基本的形式。
- 查找条件可以在不与元素的两侧进行比较的情况下确定（或使用它周围的特定元素）。
- 不需要后处理，因为每一步中，你都在检查是否找到了元素。如果到达末尾，则知道未找到该元素。



## 模板Ⅱ

```go
func binarySearch(nums []int) int {
    // 根据题目可删减
    if len(nums) == 0 {
        return -1
    }
    left, right := 0, len(nums)
    for left < right {
		mid := left + ((right - left) >> 1)
		if mid位置就是需要查找的数 {
			return mid
		} else if 如果待查找数在右侧 {
            left = mid + 1
        } else {
            // 如果待查找数在左侧
            right = mid
        }
    }
    // 处理 left == right 的情况
    if left != len(nums) && nums[left] == target {
        return left
    }
    // 没有找到该数
    return -1
}
```

 模板 #2 是二分查找的高级模板。它用于查找需要*访问数组中当前索引及其直接右邻居索引*的元素或条件。 



**关键属性**

- 一种实现二分查找的高级方法。
- 查找条件需要访问元素的直接右邻居。
- 使用元素的右邻居来确定是否满足条件，并决定是向左还是向右。
- 保证查找空间在每一步中至少有 2 个元素。
- 需要进行后处理。 当你剩下 1 个元素时，循环 / 递归结束。 需要评估剩余元素是否符合条件。





## 模板Ⅰ实战

### 二分查找

#### 题目描述

给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

**示例 1:**

```
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
```

**示例 2:**

```
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
```

**提示：**

1. 你可以假设 nums 中的所有元素是不重复的。
2. n 将在 [1, 10000]之间。
3. nums 的每个元素都将在 [-9999, 9999]之间。

#### 代码实现

直接套用模板Ⅰ

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := left + ((right - left) >> 1)
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}
```

### x的平方根

#### 题目描述

实现`int sqrt(int x)`函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

**示例 1:**

```
输入: 4
输出: 2
```

**示例 2:**

```
输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
     由于返回类型是整数，小数部分将被舍去。
```

#### 题目分析

此处需注意边界条件，也就是x为0和1的时候。其他情况则是套用模板即可。

#### 代码实现

直接套用模板Ⅰ

```go
func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    left, right := 0, x - 1
    for left <= right {
        mid := left + ((right - left) >> 1)
        tmp := mid*mid
        if tmp == x {
            return mid
        } else if tmp < x {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return right
}
```

### 猜数字大小

#### 题目描述

我们正在玩一个猜数字游戏。 游戏规则如下：
我从 **1** 到 **n** 选择一个数字。 你需要猜我选择了哪个数字。
每次你猜错了，我会告诉你这个数字是大了还是小了。
你调用一个预先定义好的接口 `guess(int num)`，它会返回 3 个可能的结果（-1，1 或 0）：

```
-1 : 我的数字比较小
 1 : 我的数字比较大
 0 : 恭喜！你猜对了！
```

**示例 :**

```
输入: n = 10, pick = 6
输出: 6
```



#### 题目分析

需要注意的是这个数字的范围是1~n，那么就意味着 `left` 和 `right` 的初值为 `1` 和 `n` 。

#### 代码实现

直接套用模板Ⅰ

```go
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    if n < 0 {
        return 1
    }
    left, right := 1, n
    for left <= right {
        mid := left + ((right - left) >> 1)
        result := guess(mid)
        if result == 0 {
            return mid
        } else if result == -1 {
            // 待猜数在左侧
            right = mid - 1
        } else {
            // 待猜数在右侧
            left = mid + 1
        }
    }
    
    return -1
}
```

### 搜索旋转排序数组

#### 题目描述

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 `[0,1,2,4,5,6,7]` 可能变为 `[4,5,6,7,0,1,2]` )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 `O(log n)` 级别。

**示例 1:**

```
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
```


**示例 2:**

```
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
```



#### 题目分析

首先，题目要求算法时间复杂度必须是`O(logn)` 这暗示需要用二分查找来做。

其次，"**假设按照升序排序的数组在预先未知的某个点上进行了旋转**"这句话很关键，意味着这个数组可能并没有发生旋转。

如果是一个**已旋转**的数组，`[4,5,6,7,0,1,2]`中，`[4,5,6,7]`称作左区间，`[0,1,2]`称作右区间。**左区间的最小元素比右区间的最大元素大**，即`4 > 2`。根据这个性质，我们可以很容易判断出一个元素是处于左区间还是右区间。例如，当前的 `mid` 所指向的元素为6，6大于右区间的最大元素2，所以它处于左区间。即一个元素如果**小于等于**数组中最后一个元素，则处于右区间，否则处于左区间。

如果是一个**未旋转**的数组，相当于 `mid` 一直处于右区间，不需要做额外的判断。



**根据可能出现的情况进行分类讨论**：

用 `m` 表示 `mid`，`t` 表示目标数

`mid` 在左区间的情况：

1. `mid` 在左区间 并且 `t` 在 `mid` 左侧

```
[左区间] [右区间]
  t m
  
mid需要往左侧移动
```

2. `mid` 在左区间 并且 `t`在左区间内 并且 `t` 在 `mid` 右侧

```
[左区间] [右区间]
   m t
mid需要往右侧移动
```

3. `mid` 在左区间 并且 `t` 在右区间

```
[左区间] [右区间]
   m       t
mid需要往右侧移动
```



`mid` 在右区间的情况：

1. `mid` 在右区间，并且 `t` 在 `mid` 右侧

```
[左区间] [右区间]
           m t
mid需要往右侧移动
```

2. `mid` 在右区间，并且 `t` 在右区间内 并且 `t` 在 `mid` 左侧

```
[左区间] [右区间]
         t m
mid需要往左侧移动
```

3. `mid` 在右区间，并且 `t` 在 `mid` 左侧

```
[左区间] [右区间]
   t       m
mid需要往左侧移动
```



实际编码的时候可以根据自己的喜好进行条件判断。此处我根据 `mid` 的移动方向作为条件判断的依据。

#### 代码实现

直接套用模板Ⅰ

```go
func search(nums []int, target int) int {
    // 注意边界条件，如果数组为空，那么last := nums[right] 这条赋值语句会报错
    if len(nums) == 0 {
        return -1
    }
    left, right := 0, len(nums) - 1
    last := nums[right] // 保存数组最后一个元素
    for left <= right {
        mid := left + ((right - left) >> 1)
        tmp := nums[mid]
        if tmp == target {
            return mid
        } else if tmp > last && target > last && target < tmp ||
        tmp <= last && !(target <= last && target > tmp){
            // 如果 mid 在左区间 并且 t 在 mid 左侧
            // 如果 mid 在右区间 并且 t 不在 mid 右侧 --> 即确保 t 在 mid 左侧
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    
    return -1
}
```



## 变体实战

### 变体一：查找第一个值等于给定值的元素

#### 题目描述

给定一个包含重复元素的有序数组，请找出第一个等于目标数的下标。

示例：

```
输入: nums = [1,3,4,5,6,8,8,8,11,18], target = 8
输出: 5
```



#### 题目分析

网上有很多关于变形二分查找的实现方法，有很多写的非常简洁，但理解起来非常烧脑，也容易写错，比如下面这个写法：

```java

public int bsearch(int[] a, int n, int value) {
  int low = 0;
  int high = n - 1;
  while (low <= high) {
    int mid = low + ((high - low) >> 1);
    if (a[mid] >= value) {
      high = mid - 1;
    } else {
      low = mid + 1;
    }
  }

  if (low < n && a[low]==value) return low;
  else return -1;
}
```

 看完这个实现之后，你是不是觉得很不好理解？如果你只是死记硬背这个写法，我敢保证，过不了几天，你就会全都忘光，再让你写，90% 的可能会写错。所以，我换了一种实现方法，你看看是不是更容易理解呢？ 

#### 代码实现

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := left + ((right - left) >> 1)
        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            // 隐含了 nums[mid] == target这个条件
            // 判断mid是否是第一个等于target的元素
            if mid == 0 || nums[mid-1] != target {
                return mid
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1
}
```



### 变体二：查找最后一个值等于给定值的元素

思路跟变体一相似，如果掌握了变体一，那么写出变体二并不难。所以就不再赘述，直接给出代码。

#### 代码实现

```go
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			// 隐含了 nums[mid] == target这个条件
			// 判断mid是否是最后一个等于target的元素
			if mid == len(nums) - 1 || nums[mid+1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
```



###  变体三：查找第一个大于等于给定值的元素 

#### 题目描述

找到第一个大于等于给定值的元素，返回其下标。

**示例：**

```
输入: nums = [3,4,6,7,10] target = 5
输出: 2
```



#### 代码实现

```go
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] >= target {
			if mid == 0 || nums[mid - 1] < target {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}

	return -1
}
```



### 变体四：查找最后一个小于等于给定值的元素

#### 代码实现

```go
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] <= target {
			if mid == len(nums) - 1 || nums[mid + 1] > target {
				return mid
			} else {
				left = mid + 1
			}
		} else {
			right = mid - 1
		}
	}

	return -1
}
```



## 总结

二分查找针对的是一个**有序**的数据集合。每次都通过跟区间的中间元素对比，将待查找的区间缩小为之前的一半，直到找到要查找的元素，或者区间被缩小为0。

### **时间复杂度**：`O(logN)`

假设数据大小是N，每次查找后数据都会缩小为原来的一半，也就是除以2。**最坏情况**下，直到查找区间被缩小为空，才停止。

被查找区间的大小变化：
$$
N, N/2, N/4, N/8, ..., N/(2^K),...1
$$
可以看出来这是一个等比数列。其中`N/(2^K) = 1`时，`K`的值就是总共缩小的次数。而每一次缩小操作只涉及两个数据的大小比较，所以，经过了`K`次区间缩小操作，时间复杂度就是`O(K)`。通过`N/(2^K) = 1`，我们可以求得`K = log2N`，所以时间复杂度就是`O(logN)`

### 二分查找的局限性

1. 二分查找依赖数组

   - 二分查找需要按照下标随机访问元素， 数组按照下标随机访问数据的时间复杂度是 O(1)，而链表随机访问的时间复杂度是 O(n)。所以，如果数据使用链表存储，二分查找的时间复杂就会变得很高。 

2.  二分查找针对的是有序数据

   - 二分查找对这一点的要求比较苛刻，数据必须是有序的。如果数据没有序，我们需要先排序。  

   - 二分查找只能用在插入、删除操作不频繁，一次排序多次查找的场景中。针对动态变化的数据集合，二分查找将不再适用。 

3.  数据量太小不适合二分查找

   - 如果要处理的数据量很小，完全没有必要用二分查找，顺序遍历就足够了。比如我们在一个大小为 10 的数组中查找一个元素，不管用二分查找还是顺序遍历，查找速度都差不多。只有数据量比较大的时候，二分查找的优势才会比较明显。

   - 有一个**例外**。如果数据之间的比较操作非常耗时，不管数据量大小，我都推荐使用二分查找。比如，数组中存储的都是长度超过 300 的字符串，如此长的两个字符串之间比对大小，就会非常耗时。我们需要尽可能地减少比较次数，而比较次数的减少会大大提高性能，这个时候二分查找就比顺序遍历更有优势。 

4.  数据量太大也不适合二分查找

   -  二分查找的底层需要依赖**数组**这种数据结构，而数组为了支持随机访问的特性，要求**内存空间连续**，对内存的要求比较苛刻。比如，我们有 1GB 大小的数据，如果希望用数组来存储，那就需要 1GB 的连续内存空间。注意这里的**“连续”**二字，也就是说，即便有 2GB 的内存空间剩余，但是如果这剩余的 2GB 内存空间都是零散的，没有连续的 1GB 大小的内存空间，那照样无法申请一个 1GB 大小的数组。而我们的二分查找是作用在数组这种数据结构之上的，所以太大的数据用数组存储就比较吃力了，也就不能用二分查找了。 

## 参考

> [一网打尽！二分查找解题模版与题型全面解析](https://www.cxyxiaowu.com/500.html)
>
> [LeetCode二分查找](https://leetcode-cn.com/explore/learn/card/binary-search/209/template-i/)