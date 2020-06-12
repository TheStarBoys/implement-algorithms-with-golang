package leetcode


func longestPalindrome005_0(s string) string {
	maxStr := ""
	max := 0
	for l := 0; l < len(s); l++ { // 外层控制l
		for r := len(s) - 1; r >= 0; r-- { // 内层控制r
			if isPalindrome005_0(s, l, r) { // 确定[l:r+1]这个区间的字符串是不是回文串
				if tmp := (r+1-l); tmp > max {
					max = tmp
					maxStr = s[l:r+1]
				}
			}
		}
	}
	return maxStr
}

func isPalindrome005_0(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func longestPalindrome005_1(s string) string {
	if s == "" {
		return ""
	}
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			// 通过这个可以减少回文判断次数
			if i - j <= r - l {
				continue
			}
			if isPalindrome005_0(s, j, i) {
				l = j
				r = i
			}
		}
	}

	return s[l:r+1]
}