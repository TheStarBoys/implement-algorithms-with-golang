package leetcode

import "math"

/*
题目：
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。
执行结果：
	用时：0 ms
	内存：2.7 MB
 */
func longestCommonPrefix(strs []string) string {
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

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}