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

func isPalindrome005_0(s string, l,r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}