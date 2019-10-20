package MI

// 题目描述：给定一个形如下述的切片，进行排序
// str := []string{"abde", "bcda", "abcd", "bcad", "bcadf", "", "bcadf"}
// 输出结果应为
// str := []string{"", "abcd", "abde", "bcad", "bcadf", "bcadf", "bcda"}

func SortStrings(str []string, indx int) []string {
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
		bucket[i] = SortStrings(bucket[i], indx+1)
	}
	// 将排序后的切片进行合并
	for i := range bucket {
		length := len(bucket[i])
		copy(str[h:h+length], bucket[i])
		h = h+length
	}
	return str
}