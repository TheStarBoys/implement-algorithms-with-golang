package leetcode

import "strings"
// 自己实现的算法，效率极低
func myAtoi008_0(str string) int {
	str = strings.TrimSpace(str)
	if str == "" || str[0] != '+' && str[0] != '-' && str[0] < '0' || str[0] > '9' {
		return 0
	}
	numStr := ""
	flag := ""
	if str[0] == '+' || str[0] == '-' {
		flag = string(str[0])
		numStr = str[1:]
	}else {
		numStr = str
	}
	num := 0
	for i := 0; i < len(numStr); i++ {
		if numStr[i] < '0' || numStr[i] > '9' {
			break
		}
		curNum := 0
		switch numStr[i] {
		case '1':
			curNum = 1
		case '2':
			curNum = 2
		case '3':
			curNum = 3
		case '4':
			curNum = 4
		case '5':
			curNum = 5
		case '6':
			curNum = 6
		case '7':
			curNum = 7
		case '8':
			curNum = 8
		case '9':
			curNum = 9
		}
		num = num * 10 + curNum
		if num > 1 << 31 - 1 {
			break
		}
	}
	if flag == "-" {
		num =  -num
	}
	if num > 1 << 31 - 1 {
		return 1 << 31 - 1
	}else if num < -(1 << 31)  {
		return -(1 << 31)
	}
	return num
}
