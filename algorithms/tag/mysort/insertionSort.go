package mysort

// 插入排序
// 将排序数组a想象为牌堆，默认抽第一张牌（下标为0）放入手牌（从0 ~ j均为手牌），手牌是已排序的
// 从 j = 1 开始摸牌，每抽到一张牌（记为 key = a[j]），跟手牌比较寻找插入位置
// 若 key 比当前比较的手牌 a[i] 大，则将手牌往后挪一位
// 找到插入位置后，key放置到空位 a[i + 1]
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