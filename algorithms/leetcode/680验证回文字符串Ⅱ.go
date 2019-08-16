package leetcode

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1{
		if s[i] != s[j] {
			return isPalindrome1(s, i, j-1) || isPalindrome1(s, i+1, j)
		}
	}
	return true
}

func isPalindrome1(s string, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func ValidPalindrome(s string) bool {
	return validPalindrome(s)
}