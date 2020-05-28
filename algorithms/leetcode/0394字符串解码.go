package leetcode

import (
	"strconv"
	"strings"
)

// 辅助栈
func decodeString0394_0(s string) string {
	res := ""
	multi := 0
	stack_multi := make([]int, 0)
	stack_res := make([]string, 0)

	for i := range s {
		if s[i] == '[' {
			stack_multi = append(stack_multi, multi)
			stack_res = append(stack_res, res)
			multi = 0
			res = ""
		} else if s[i] == ']' {
			tmp := ""
			cur_multi := stack_multi[len(stack_multi) - 1]
			stack_multi = stack_multi[:len(stack_multi) - 1]
			for i := 0; i < cur_multi; i++ { tmp += res }
			res = stack_res[len(stack_res) - 1] + tmp
			stack_res = stack_res[:len(stack_res) - 1]
		} else if '0' <= s[i] && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			multi = multi * 10 + num
		} else { res += string(s[i]) }
	}
	return res
}

// 递归
func decodeString_0394_1(s string) string {
	if s == "" {
		return ""
	}
	left := strings.Index(s, "[")
	if left == -1 {
		return s
	}
	// 去找形如 "a2[c]" 中第一个非字符下标
	charIndx := 0
	for ; charIndx < len(s) && s[charIndx] <= 'z' && s[charIndx] >= 'a'; charIndx++ {}
	if charIndx != 0 {
		// 将charIndx后的字符串解码后 进行拼接
		return s[:charIndx] + decodeString_0394_1(s[charIndx:])
	}
	k, _ := strconv.Atoi(s[:left])
	// 找到该左括号对应匹配的右括号
	count := 1
	right := 0
	for i := left+1; i < len(s); i++ {
		if s[i] == '[' {
			count++
		}else if s[i] == ']' {
			count--
			if count == 0 {
				right = i
				break
			}
		}
	}
	// 确保括号中的内容是解码过的
	str := decodeString_0394_1(s[left+1:right])
	// 将字符串进行重复
	tmp := str
	for i := 0; i < k-1; i++ {
		str += tmp
	}
	// 如果对应的右括号后仍有字符，将后续字符继续解码
	if right < len(s) {
		return str + decodeString_0394_1(s[right+1:])
	}

	return str
}