package main

import (
	"fmt"
	"strings"
)
func main() {
	fmt.Println(fun(5))
}
func fun(n int) int {
	if n == 1 {
		return 1
	}
	return n * fun(n - 1)
}
func delete(l, r int, s string) int {
	count := 0
	for ; l < r; l, r = l + 1, r - 1 {
		if s[l] != s[r] {
			count = delete(l+1, r, s) + 1
			tmp := delete(l, r-1, s) + 1
			if count > tmp {
				count = tmp
			}
		}
	}
	return count
}

func myAtoi(str string) int {
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
			return 1 << 31 - 1
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