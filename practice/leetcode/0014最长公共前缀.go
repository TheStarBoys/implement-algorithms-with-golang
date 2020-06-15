package leetcode

import (
	"math"
	"strings"
)

/*
题目：
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。
执行结果：
	用时：0 ms
	内存：2.7 MB
 */
func longestCommonPrefix014_0(strs []string) string {
	s := ""
	if len(strs) == 0 {
		return s
	}
	temp := [][]byte{}
	// 算出最小的列长度
	minLen := math.MaxInt32
	for _, str := range strs {
		temp = append(temp, []byte(str))
		if cur := len([]byte(str)); minLen > cur {	// 如果最小行长度 > 当前行的长度
			minLen = cur
		}
	}
	for j := 0; j < minLen; j++ {	// j控制列
		for i := 0; i < len(temp); i++ { // i控制行
			if i < len(temp) -1 && temp[i][j] != temp[i+1][j] { // j列的元素跟所有行的比较
				return s
			}
			if i == len(temp) -1 {
				s += string(temp[i][j])
			}
		}
	}
	return s
}

func longestCommonPrefix014_1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	for i := 1; i < len(strs); i ++ { // find minLen
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
	sumStr := ""
	for i := 0; i < minLen; i++ { // i control col
		curStr := strs[0][i]
		for j := 1; j < len(strs); j++ { // j control row
			if strs[j][i] != curStr {
				return sumStr
			}
		}
		sumStr += string(curStr)
	}
	return sumStr
}
// 水平扫描法
func longestCommonPrefix014_2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix) - 1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}
// 分治法
func longestCommonPrefix014_3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return commonPrefix014_3(strs)
}
func commonPrefix014_3(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}else if len(strs) == 2 {
		prefix := strs[0]
		for strings.Index(strs[1], prefix) != 0 {
			prefix = prefix[:len(prefix) - 1]
		}
		return prefix
	}

	mid := len(strs) / 2
	lcpLeft := commonPrefix014_3(strs[:mid])
	lcpRight := commonPrefix014_3(strs[mid:])
	return commonPrefix014_3([]string{lcpLeft, lcpRight})
}
// 分治法效率高的解
func longestCommonPrefix014_4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return longCommonPrefix014_4(strs, 0, len(strs)-1)
}
func longCommonPrefix014_4(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}else {
		mid := (l + r) / 2
		lcpLeft := longCommonPrefix014_4(strs, l, mid)
		lcpRight := longCommonPrefix014_4(strs, mid + 1, r)
		return commonPrefix014_4(lcpLeft, lcpRight)
	}
}
func commonPrefix014_4(left, right string) string {
	min := len(left)
	if len(right) < min {
		min = len(right)
	}
	for i := 0; i < min; i++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:min]
}
// 二分法
func longestCommonPrefix014_5(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
	l, h := 0, minLen - 1
	for l <= h {
		mid := l + (h - l ) / 2
		if isCommonPrefix014_5(l, mid, strs) {
			l = mid + 1
		}else {
			h = mid - 1
		}
	}
	return strs[0][:h + 1]
}
func isCommonPrefix014_5(l, h int, strs []string) bool {
	prefix := strs[0][l : h + 1]
	for i := 1; i < len(strs); i++ {
		if strings.Index(strs[i][l : h + 1], prefix) == -1 {
			return false
		}
	}
	return true
}

// 二分查找，或许是更易懂的形式
func longestCommonPrefix014_6(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
	l, h := 0, minLen - 1
	mid := l + ((h - l) >> 1)
	if isCommonPrefix014_5(l, mid, strs) {
		if mid == minLen - 1 || isCommonPrefix014_5(l, mid+1, strs) == false {
			return strs[0][:mid+1]
		} else {
			l = mid + 1
		}
	} else {
		h = mid - 1
	}
	return ""
}