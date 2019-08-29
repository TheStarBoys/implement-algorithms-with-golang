package leetcode
/*
题目:
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
进阶：
	你能不将整数转为字符串来解决这个问题吗
 */
func isPalindrome009_0(x int) bool {
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

// 官方解法：
func isPalindrome009_1(x int) bool {
	/*
	除了负数外，如果一个数的最后一位是0，他的第一位也必须是0，不存在0120这样的写法，所以直接返回false
	 */
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	reverseNum := 0
	/*
	当reverseNum > x的时候，说明这个数翻转到了这个数的中间位置了
	 */
	for x > reverseNum {
		reverseNum = reverseNum * 10 + x % 10
		x /= 10
	}
	/*
	直接判断x == reverseNum，如果不相等，那么这个数的位数一定是个奇数
	类似于：12321， x = 12, reverseNum = 123, 直接去除掉最后一位就可以了
	 */
	return x == reverseNum || x == reverseNum / 10
}