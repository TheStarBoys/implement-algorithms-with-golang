package leetcode

// 单调栈，逆序遍历
func dailyTemperatures739_0(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0) // 单调栈，非递增排序

	for i := len(T) - 1; i >= 0; i-- {
		// 当前元素比栈顶元素大， 出栈， 重新调整栈直至满足要求
		for len(stack) != 0 && T[i] >= T[stack[len(stack) - 1]] {
			stack = stack[:len(stack) - 1]
		}
		// 栈为空， 即后面没有比当前天温度高的
		if len(stack) == 0 {
			res[i] = 0
		} else { // 不为空，栈顶元素对应的下标减去当前下标即为经过几天后温度比当前天温度高
			res[i] = stack[len(stack) - 1] - i
		}
		// 当前元素进栈
		stack = append(stack, i)
	}
	return res

}
