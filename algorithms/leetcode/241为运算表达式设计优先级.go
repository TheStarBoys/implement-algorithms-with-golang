package leetcode

import "strconv"

// 解题思路：
// 分治思想：将一个大问题，分解成若干个小问题，满足最优子结构；
// 在问题规模为没有运算符的之后，直接将字符串赋值给out

// 根据运算符，分割运算符左右两边的表达式，通过递归求得左右表达式的值
// 再根据当前运算符，计算得到运算结果
func diffWaysToCompute(s string) []int {
	out := []int{}
	for i, b := range []byte(s) {
		if b != '+' && b != '-' && b != '*' {
			continue
		}
		sll := diffWaysToCompute(s[:i])
		slr := diffWaysToCompute(s[i+1:])
		for _, il := range sll {
			for _, ir := range slr {
				switch b {
				case '+':
					out = append(out, il+ir)
				case '-':
					out = append(out, il-ir)
				case '*':
					out = append(out, il*ir)
				}
			}
		}
	}
	if 0 == len(out) { // 当前字符串没有运算符，则直接把字符串的值赋给out
		i, _ := strconv.Atoi(s)
		out = []int{i}
	}
	return out
}

