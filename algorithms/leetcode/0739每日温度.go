package leetcode

// 单调栈，逆序遍历
func dailyTemperatures0739_0(T []int) []int {
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

// 单调栈，正序遍历
func dailyTemperatures0739_1(T []int) []int {
	if len(T) == 0 {
		return []int{}
	}

	ans := make([]int, len(T))
	stack := make([]int, 0, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			prevIndx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndx] = i - prevIndx
		}
		stack = append(stack, i)
	}

	return ans
}

// 逆序遍历
func dailyTemperatures0739_2(T []int) []int {
	if len(T) == 0 {
		return []int{}
	}
	var (
		n = len(T)
		res = make([]int, n)
		maxVal = T[n-1]
		maxIndx = n - 1
	)

	for i := len(T) - 2; i >= 0; i-- {
		if maxVal <= T[i] {
			maxVal = T[i]
			maxIndx = i
		}
		if T[i] >= maxVal {
			continue
		}
		day := 1
		for j := i+1; j <= maxIndx; j++ {
			if T[i] < T[j] {
				res[i] = day
				break
			}
			day++
		}
	}

	return res
}