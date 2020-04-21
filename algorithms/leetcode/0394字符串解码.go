package leetcode

import "strconv"

// 辅助栈
func decodeString394_0(s string) string {
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
