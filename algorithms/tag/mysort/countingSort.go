package mysort

// 计数排序：
// 排序数组需要是非负整数，若不是，需要进行一定的转换
// 只能用在数据范围不大的场景中，如果数据范围k比要排序的数据n大很多
// 就不适合用基数排序了。

// 假设数组中存储的都是非负整数
func countingSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	// 查找数组中数据的范围
	max := a[0]
	for i := 1; i < n; i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	// 申请一个计数数组c，下标大小[0, max]
	c := make([]int, max + 1)

	// 计算每个元素的个数，放入c中
	for i := 0; i < n; i++ {
		c[a[i]]++
	}

	// 依次累加
	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	// 临时数组r，存储排序后的结果
	r := make([]int, n)
	// 计算排序的关键步骤，有点难理解
	for i := n-1; i >= 0; i-- {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}

	// 将结果拷贝给a数组
	copy(a, r)
}

func countingSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 找出数组中最大值
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	// 初始化计数数组
	countArr := make([]int, max+1)

	// 计数
	for i := 0; i < len(arr); i++ {
		countArr[arr[i]]++
		arr[i] = 0
	}

	// 排序
	index := 0
	for i := 0; i < len(countArr); i++ {
		for countArr[i] > 0 {
			arr[index] = i
			index++
			countArr[i]--
		}
	}
}