package mysort

import "fmt"

// 桶排序，需要明确输入数组的取值范围
// 时间O(M+N) 空间O(M) M是桶的长度, 如果输出需要是切片的话，空间为O(M+N)
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