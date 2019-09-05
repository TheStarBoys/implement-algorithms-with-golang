package leetcode

import "strconv"

func evalRPN150_0(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		if v == "+" {
			num2 := stack[len(stack) - 1]
			num1 := stack[len(stack) - 2]
			stack = stack[:len(stack) - 2]
			stack = append(stack, num1 + num2)
		} else if v == "-" {
			num2 := stack[len(stack) - 1]
			num1 := stack[len(stack) - 2]
			stack = stack[:len(stack) - 2]
			stack = append(stack, num1 - num2)
		} else if v == "*" {
			num2 := stack[len(stack) - 1]
			num1 := stack[len(stack) - 2]
			stack = stack[:len(stack) - 2]
			stack = append(stack, num1 * num2)
		} else if v == "/" {
			num2 := stack[len(stack) - 1]
			num1 := stack[len(stack) - 2]
			stack = stack[:len(stack) - 2]
			stack = append(stack, num1 / num2)
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[len(stack) - 1]
}
