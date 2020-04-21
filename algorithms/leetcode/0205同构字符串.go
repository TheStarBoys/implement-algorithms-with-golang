package leetcode

// 待优化
func isIsomorphic205_0(s string, t string) bool {
	if len(s) != len(t) { return false }
	mappings := make(map[byte]byte) // 完成映射
	set := make(map[byte]bool) // 不允许两个字符映射向同一个字符
	for i := range s {
		if v, ok := mappings[s[i]]; ok {
			if v != t[i] { return false } // 不满足映射条件
			continue
		}
		if set[t[i]] { return false }
		set[t[i]] = true
		mappings[s[i]] = t[i]
	}
	return true
}

// 通过一个辅助函数
func isIsomorphic205_1(s string, t string) bool {
	if len(s) != len(t) { return false }
	return helper205_1(s, t) && helper205_1(t, s)
}
func helper205_1(s, t string) bool {
	mappings := make(map[byte]byte)
	for i := range s {
		if v, ok := mappings[s[i]]; ok {
			if v != t[i] { return false } // 不满足映射条件
			continue
		}
		mappings[s[i]] = t[i]
	}
	return true
}
// 将字符串转成[]rune效率更高，但是空间复杂度会增加
func isIsomorphic205_2(s string, t string) bool {
	if len(s) != len(t) { return false }
	return helper205_2([]rune(s), []rune(t)) && helper205_2([]rune(t), []rune(s))
}
func helper205_2(s, t []rune) bool {
	mappings := make(map[rune]rune)
	for i := range s {
		if v, ok := mappings[s[i]]; ok {
			if v != t[i] { return false } // 不满足映射条件
			continue
		}
		mappings[s[i]] = t[i]
	}
	return true
}