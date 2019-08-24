package leetcode
// 递归
func reverseString(s []byte)  {
	helper_reverseString(0, len(s)-1, s)
}
func helper_reverseString(i, j int, s []byte) {
	if i >= j {
		return
	}
	s[i], s[j] = s[j], s[i]
	helper_reverseString(i+1, j-1, s)
}
