package mysort

// 插入排序
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