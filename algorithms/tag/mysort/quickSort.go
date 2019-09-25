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
