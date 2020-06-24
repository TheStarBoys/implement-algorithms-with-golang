package MI

// 题目描述：给定一个形如下述的切片，进行排序
// str := []string{"abde", "bcda", "abcd", "bcad", "bcadf", "", "bcadf"}
// 输出结果应为
// str := []string{"", "abcd", "abde", "bcad", "bcadf", "bcadf", "bcda"}

// 效率较高，空间浪费大
func SortStrings001_0(str []string, indx int) []string {
	// 'a' -> 'z' 26个字母，映射到26个桶中
	bucket := make([][]string, 26)
	h := 0 // 合并时的起始下标
	for i := 0; i < len(str); i++ {
		if indx >= len(str[i]) { // 当前元素比较短
			// 插入排序，将短的，插到前面合适的位置
			j := i-1
			key := str[i]
			for ; j >= 0 && len(str[j]) > len(key); j-- {
				str[j+1] = str[j]
			}
			// 找到的j的长度 <= i
			str[j+1] = key
			h++
			continue
		}
		site := str[i][indx]-'a'
		bucket[site] = append(bucket[site], str[i])
	}
	for i := range bucket {
		// 如果桶中的元素小于1，那么就不需要排序了
		if len(bucket[i]) <= 1 {
			continue
		}
		// 递归的将桶里的数据也排好序
		bucket[i] = SortStrings001_0(bucket[i], indx+1)
	}
	// 将排序后的切片进行合并
	for i := range bucket {
		length := len(bucket[i])
		copy(str[h:h+length], bucket[i])
		h = h+length
	}
	return str
}

// 新思路
func SortStrings001_1(strs []string) []string {
	for i := 0; i < len(strs) - 1; i++ {
		for j := 1; j < len(strs) - i; j++ {
			if isSmall001_1(strs[j-1], strs[j]) == false {
				strs[j-1], strs[j] = strs[j], strs[j-1]
			}
		}
	}

	return strs
}

// 判断字符串 a 是否小于 b
func isSmall001_1(a, b string) bool {
	var i, j int

	if a == "" {
		return true
	} else if b == "" {
		return false
	}

	for ; i < len(a) && j < len(b) ; i, j = i+1, j+1 {
		// 'a' - 'b' = -1

		// 以下注释代码有问题是由于 a[i] 为 byte类型变量
		// 底层为uint8，如果类似于 a[i] - b[j] = 0 - 2 的情况
		// 结果并不为-2
		// (0000 0000) - (0000 0010)
		// (0000 0000) + (1111 1110)
		// = 254

		// 如果是 int8 则会是
		// (0000 0000) - (0000 0010)
		// (0000 0000) + (1000 0010)
		// = -2

		//if tmp := a[i] - b[j]; tmp < 0 {
		if tmp := int(a[i]) - int(b[j]); tmp < 0 {
			return true
		} else if tmp > 0 {
			return false
		}
	}

	if i >= len(a) {
		return true
	}

	return false
}

// 归并排序
func SortStrings001_2(strs []string) {
	if len(strs) < 2 { return }
	mid := len(strs) >> 1
	SortStrings001_2(strs[:mid])
	SortStrings001_2(strs[mid:])
	merge001_2(strs)
}

func merge001_2(strs []string) {
	mid := len(strs) >> 1
	tmp := make([]string, len(strs))
	left := strs[:mid]
	right := strs[mid:]
	indx := 0
	i, j := 0, 0
	for ; i < len(left) && j < len(right); indx++ {
		if left[i] <= right[j] {
			tmp[indx] = left[i]
			i++
		} else {
			tmp[indx] = right[j]
			j++
		}
	}
	if i == len(left) {
		copy(tmp[indx:], right[j:])
	}
	if j == len(right) {
		copy(tmp[indx:], left[i:])
	}
	copy(strs, tmp)
}