package leetcode


// 387 字符串中的第一个唯一字符
func firstUniqChar387_0(s string) int {
	table := make(map[byte]int)
	for i := range s {
		table[s[i]]++
	}
	for i := range s {
		if  table[s[i]] == 1 {
			return i
		}
	}
	return -1
}

