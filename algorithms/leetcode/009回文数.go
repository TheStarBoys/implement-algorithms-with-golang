package leetcode
/*
题目:
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
进阶：
	你能不将整数转为字符串来解决这个问题吗
 */
func isPalindrome(x int) bool {
	if x < 0 {	// 负数不是回文数
		return false
	}
	y, z := 0, x
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	if y == z {
		return true
	}
	return false
}
