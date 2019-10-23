package leetcode

// 累减
func divide029_0(dividend int, divisor int) int {
	flag := 1
	if divisor < 0 && dividend > 0 || divisor > 0 && dividend < 0 {
		flag = -1
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}
	count := 0
	for dividend > 0 {
		dividend -= divisor
		count++
	}
	return flag * (count-1)
}

