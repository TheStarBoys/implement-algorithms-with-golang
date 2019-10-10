package leetcode

import "fmt"

// 方法思路：每次只找个位数是否是0，该方法会在阶乘过大的时候溢出
func trailingZeroes172_0(n int) int {
	num := n
	for i := n-1; i >= 2; i-- {
		num *= i
	}
	fmt.Println("num= ", num)
	count := 0
	for num != 0 {
		if num % 10 == 0 {
			count++
		} else {
			break
		}
		num /= 10
	}
	return count
}


// 将寻找零的个数，转换为寻找5的个数
func trailingZeroes172_1(n int) int {
	// 题解链接：
	// https://leetcode-cn.com/problems/factorial-trailing-zeroes/solution/ji-suan-cheng-shu-zhong-5de-shu-liang-by-user8208/
	sum := n / 5
	m := n
	for m / 5 != 0 && m >= 5 {
		m /= 5
		sum += m / 5
	}
	return sum
}
