package leetcode

// 125 验证回文串, 因为字符串的拼接时间复杂度非常大，换成bytes后效率从340ms到4ms
func isPalindrome125_0(s string) bool {
	if len(s) <= 1 {
		return true
	}
	b := []byte{}
	for i := range s {
		if 65 <= s[i] && s[i] <= 90 {
			b = append(b, s[i] + 32) // 'A' --> 'a'
			continue
		}
		if '0' <= s[i] && s[i] <= '9' || 'a' <= s[i] && s[i] <= 'z' {
			b = append(b, s[i])
		}
	}
	l, r := 0, len(b) - 1
	if r < 0 { return true }
	for l < r {
		if b[l] != b[r] {
			return false
		}
		l++
		r--
	}
	return true
}
// 优化方案，用变长的[]byte来代替string
func isPalindrome125_1(s string) bool {
	if len(s) <= 1 {
		return true
	}
	str :=""
	for i := range s {
		if 65 <= s[i] && s[i] <= 90 {
			str += string(s[i] + 32) // 'A' --> 'a'
			continue
		}
		if '0' <= s[i] && s[i] <= '9' || 'a' <= s[i] && s[i] <= 'z' {
			str += string(s[i])
		}
	}
	l, r := 0, len(str) - 1
	if r < 0 { return true }
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
