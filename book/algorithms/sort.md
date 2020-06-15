# 排序算法

## 如何分析一个排序算法

### 排序算法的执行效率

**1. 最好情况、最坏情况、平均情况时间复杂度**

**2. 时间复杂度的系数、常数、低阶**

**3. 比较次数和交换（或移动）次数**

### 排序算法的内存消耗

### 排序算法的稳定性

## 冒泡排序

时间复杂度：`O(N)`

空间复杂度：`O(1)`

内层循环每次两两比较，将最大值向右移到排序后的正确位置 `len(nums) - i`

由于内层循环每次将当前最大值放在最后，当外层循环到 `len(nums) - 1` 时，剩下最后一个元素已经为正确位置，所以该元素不需要额外排序

```go
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		for j := 1; j < len(nums) - i; j++ {
			if nums[j - 1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}
```



## 桶排序

时间复杂度：`O(M+N)`

空间复杂度：`O(M)`，如果需要输出是切片的话，空间为`O(M+N)`

`M`是桶的大小

```go
// 输入在0~1000之间

// 桶排序需要知道一个数的取值范围，再根据取值范围生成相应数量的桶
// 遍历nums，每取出一个数在对应的桶中放一个旗子
// 最后遍历桶，将得到正确排序
func bucketSort(nums []int) {
	bucket := make([]int, 1001)
	for _, v := range nums {
		bucket[v]++
	}
	for i := range bucket {
		for j := 0; j < bucket[i]; j++ {
			fmt.Printf("%d ", i)
		}
	}
}
```

时间复杂度：`O(n*log(n/m))`

空间复杂度：`O(m)`


```go
// 如果要排序的数据有 n 个，我们把它们均匀地划分到 m 个桶内，每个桶里就有 k=n/m 个元素。
// 每个桶内部使用快速排序，时间复杂度为 O(k * logk)。m 个桶排序的时间复杂度就是 O(m * k * logk)，因为 k=n/m
// 所以整个桶排序的时间复杂度就是 O(n*log(n/m))。
// 当桶的个数 m 接近数据个数 n 时，log(n/m) 就是一个非常小的常量，这个时候桶排序的时间复杂度接近 O(n)
func bucketSort2(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	// 找到最大值
	max := arr[0]
	for i := 1; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	buckets := make([][]int, n)

	for i := 0; i < n; i++ {
		// arr[i] 的取值范围：[0, max]
		// arr[i] * (n-1) 的取值范围：[0, max*(n-1)]
		// arr[i] * (n - 1) / max 的取值范围：[0, n-1]
		index := arr[i] * (n - 1) / max // 桶序号
		buckets[index] = append(buckets[index], arr[i]) // 加入到对应的桶中
	}

	index := 0 // 标记数组位置
	for i := 0; i < n; i++ {
		if len(buckets[i]) > 0 {
			quickSort2(buckets[i], 0, len(buckets[i])-1) // 桶内快排
			copy(arr[index:], buckets[i])
			index += len(buckets[i])
		}
	}
}
```



## 插入排序

时间复杂度：

空间复杂度：

```go
// 将排序数组a想象为牌堆，默认抽第一张牌（下标为0）放入手牌（从0 ~ j均为手牌），手牌是已排序的
// 从 j = 1 开始摸牌，每抽到一张牌（记为 key = a[j]），跟手牌比较寻找插入位置
// 若 key 比当前比较的手牌 a[i] 大，则将手牌往后挪一位
// 找到插入位置后，key放置到空位 a[i + 1]

// 升序排序
func InsertionSort_Up(a []int) {
	for j := 1; j < len(a); j++ {
		i := j - 1
		key := a[j]
		for i >= 0 && a[i] > key {
			a[i + 1] = a[i]
			i--
		}
		a[i + 1] = key
	}
}
// 降序排序
func InsertionSort_Down(a []int) {
	for j := 1; j < len(a); j++ {
		i := j - 1
		key := a[j]
		for i >= 0 && a[i] < key {
			a[i + 1] = a[i]
			i--
		}
		a[i + 1] = key
	}
}
```



## 希尔排序

```shell
func shellSort(nums []int) {
	gap := 1
	// 区间取值没有定论，大概是在1/2 到 1/3之间
	for gap < len(nums) {
		gap = gap * 3 + 1 // 1 -> 4 -> 13
	}

	for gap > 0 {
		// 当 gap == 1时，就是插入排序
		// 通过gap来跨区间，插排
		for i := gap; i < len(nums); i++ { // 6
			tmp := nums[i] // 4
			j := i - gap // 2
			for j >= 0 && nums[j] > tmp {
				nums[j + gap] = nums[j]
				j -= gap
			}
			nums[j + gap] = tmp
		}
		gap = int(math.Floor(float64(gap) / 3))
	}
}
```



## 归并排序

```go
func mergeSort(A []int, p, r int) {
	if p < r {
		q := int(math.Floor(float64((p + r) / 2)))
		mergeSort(A, p, q)
		mergeSort(A, q + 1, r)
		merge(A, p, q, r)
	}
}
func merge(A []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left, right := make([]int, n1 + 1), make([]int, n2 + 1)
	copy(left, A[p:q + 1])
	copy(right, A[q + 1:r + 1])
	left[n1] = math.MaxInt64
	right[n2] = math.MaxInt64
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			A[k] = left[i]
			i++
		} else {
			A[k] = right[j]
			j++
		}
	}
}
```



## 快速排序

时间复杂度：

空间复杂度：

```go
func quickSort(arr []int, left, right int) {
	var i, j, temp int
	if left > right { return }

	temp = arr[left] // temp就是基准数
	i, j = left, right
	for i != j {
		// 顺序很重要，先从右往左找。找到比基准数小的数
		for ; arr[j] >= temp && i < j; j-- {}
		// 从左往右找，找到比基准数大的数
		for ; arr[i] <= temp && i < j; i++ {}
		if i < j { // 哨兵i和哨兵j没有相遇时
			arr[i], arr[j] = arr[j], arr[i] // 交换位置
		}
	}
	// 基准数归位
	arr[left] = arr[i]
	arr[i] = temp

	quickSort(arr, left, i-1) // 继续处理左边的，这是一个递归的过程
	quickSort(arr, i+1, right) // 继续处理右边的，这是一个递归的过程
}

```



## 引用

> [这或许是东半球分析十大排序算法最好的一篇文章](https://www.cxyxiaowu.com/725.html)

