package leetcode

// 方法1：寻找“水池”的左右端点
func trap0042_0(height []int) int {
	var ans int
	// l 水池的左端
	for l := 0; l < len(height) - 2; {
		// 找到水池的右端
		r := l+1 // 这个右端是满足 height[r] >= height[l] 的最佳情形
		maxIndx := r // 如果右端所有情况均比 左端矮，就取其中最高的
		for ; r < len(height) && height[r] < height[l]; r++ {
			if height[maxIndx] < height[r] {
				maxIndx = r
			}
		}
		if r == len(height) { // 取右侧最高的作为水池的右端
			ans += calcuPool0042_0(height[l:maxIndx+1])
			l = maxIndx
		} else { // 找到了满足的最佳右端
			ans += calcuPool0042_0(height[l:r+1])
			l = r
		}
	}
	return ans
}

// 计算水池的最大盛水
func calcuPool0042_0(height []int) int {
	if len(height) < 3 {
		return 0
	}
	n := len(height)
	// 最大水平面
	maxLevel := min(height[0], height[n-1])
	// 宽度
	width := n-2
	// 最大容积
	maxVolume := width * maxLevel
	// 最大容积减去水池中不可盛水的空间
	for i := 1; i < n-1; i++ {
		maxVolume -= min(height[i], maxLevel)
	}

	return maxVolume
}

// LeetCode官方解法1：直观法
func trap0042_1(height []int) int {
	var ans int
	for i := 0; i < len(height); i++ {
		max_left, max_right := 0, 0
		for j := i; j >= 0; j-- {
			max_left = max(max_left, height[j])
		}
		for j := i; j < len(height); j++ {
			max_right = max(max_right, height[j])
		}
		ans += min(max_left, max_right) - height[i]
	}
	return ans
}

// LeetCode官方解法2：动态编程
func trap0042_2(height []int) int {
	var ans int
	n := len(height)
	if n == 0 { return 0 }
	left_max := make([]int, n)
	right_max := make([]int, n)
	// 找到数组中从下标 i 到最左端最高的条形块高度 left_max。
	left_max[0] = height[0]
	for i := 1; i < n; i++ {
		left_max[i] = max(left_max[i-1], height[i])
	}
	// 找到数组中从下标 i 到最右端最高的条形块高度 \text{right\_max}right_max
	right_max[n-1] = height[n-1]
	for i := n-2; i >= 0; i-- {
		right_max[i] = max(right_max[i+1], height[i])
	}
	// 遍历数组累加答案
	for i := 1; i < n; i++ {
		ans += min(left_max[i], right_max[i]) - height[i]
	}
	return ans
}

// LeetCode官方解法3：单调递减栈
func trap0042_3(height []int) int {
	var ans int
	n := len(height)
	if n == 0 { return 0 }
	stack := make([]int, 0, n)
	stackTop := func() int {
		return stack[len(stack)-1]
	}
	for i := 0; i < n; i++ {
		for len(stack) != 0 && height[i] >= height[stackTop()] {
			top := stackTop()
			// pop
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			// 宽度
			distance := i - stackTop() - 1
			// 高度
			bounded_height := min(height[stackTop()], height[i]) - height[top]
			ans += bounded_height * distance
		}
		stack = append(stack, i)
	}

	return ans
}

// LeetCode官方解法4：双指针
func trap0042_4(height []int) int {
	var ans int
	l, r := 0, len(height) - 1
	left_max, right_max := 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] > left_max {
				left_max = height[l]
			} else {
				ans += left_max - height[l]
			}
			l++
		} else {
			if height[r] > right_max {
				right_max = height[r]
			} else {
				ans += right_max - height[r]
			}
			r--
		}
	}

	return ans
}