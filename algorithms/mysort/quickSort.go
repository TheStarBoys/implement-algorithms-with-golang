package mysort

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

// quick sort 2
func quickSort2(arr []int, p, r int) {
	if p >= r { return }

	// 分区点将数组分为了三个部分
	// 左侧的值均小于分区点值
	// 右侧的值均大于分区点值
	q := partition(arr, p, r) // 获取分区点
	quickSort2(arr, p, q-1)
	quickSort2(arr, q+1, r)
}

// 我们通过游标 i 把 A[p…r-1]分成两部分。A[p…i-1]的元素都是小于 pivot 的，
// 我们暂且叫它“已处理区间”，A[i…r-1]是“未处理区间”。
// 我们每次都从未处理的区间 A[i…r-1]中取一个元素 A[j]，与 pivot 对比，
// 如果小于 pivot，则将其加入到已处理区间的尾部，也就是 A[i]的位置。
func partition(arr []int, p, r int) int {
	i := p
	pivot := arr[r] // 默认选最后一个元素作为分区点
	for j := p; j < r; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]

	return i
}


// 如果快速排序的数组本身就是有序的或者接近有序的
// 那么每次选择最后一个元素作为分区点，那么快速排序就会变得非常糟糕
// 如何优化快速排序？

// 1. 三数取中法
// 在数组的首、中、尾取三个数取中间值作为分区点。如果数组很大，可能就需要“五数取中”、“八数取中”了

// 2. 随机法
// 随机选取一个数组的元素，平均下来分区点相对合适，出现退化为O(n^2)的几率不大