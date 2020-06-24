package leetcode

// 125 验证回文串, 因为字符串的拼接时间复杂度非常大，换成bytes后效率从340ms到4ms
func isPalindrome0125_0(s string) bool {
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

// 优化方案，用变长的[]byte来代替string
func isPalindrome0125_1(s string) bool {
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

// 双指针一次遍历
func isPalindrome0125_2(s string) bool {
	l, r := 0, len(s) - 1
	for l < r {
		if !isValidChar0125_2(s[l]) {
			l++
			continue
		}
		if !isValidChar0125_2(s[r]) {
			r--
			continue
		}
		if !toLowerCmpr0125_2(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isValidChar0125_2(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' ||
		b >= '0' && b <= '9' {
		return true
	}
	return false
}

func toLowerCmpr0125_2(b1, b2 byte) bool {
	if b1 >= 'A' && b1 <= 'Z' {
		b1 += 32
	}
	if b2 >= 'A' && b2 <= 'Z' {
		b2 += 32
	}

	return b1 == b2
}