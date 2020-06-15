package dynamicProgramming

// items: 物品重量，w 背包最大承重
// 回溯做法
func knapsack1(items []int, w int) {
	maxW = 0
	knapsack_backtrack1(items, 0, 0, w)
}

// items: 物品重量，curi 当前抉择物品，curW 背包当前重量，w 背包最大承重
var (
	maxW int // 结果
	mem [][]bool
)
func knapsack_backtrack1(items []int, curi, curW, w int) {
	if curi >= len(items) { // 回溯的终止条件
		if curW > maxW { // 满足条件，则更新结果
			maxW = curW
		}
		return
	}
	// 不选 curi 物品
	knapsack_backtrack1(items, curi+1, curW, w)
	// 满足背包承重的情况下，选择 curi 物品
	if items[curi] + curW <= w {
		knapsack_backtrack1(items, curi+1, curW + items[curi], w)
	}
}

// 带备忘录的回溯
func knapsack2(items []int, w int) {
	// 初始化备忘录
	maxW = 0
	mem = make([][]bool, len(items))
	for i := 0; i < len(mem); i++ {
		mem[i] = append(mem[i], make([]bool, w+1)...)
	}
	knapsack_backtrack2(items, 0, 0, w)
}

func knapsack_backtrack2(items []int, curi, curW, w int) {
	if curi >= len(items) {
		if curW > maxW {
			maxW = curW
		}
		return
	}
	if mem[curi][curW] { return } // 有重复计算，直接返回
	mem[curi][curW] = true
	knapsack_backtrack2(items, curi+1, curW, w)
	if items[curi] + curW <= w {
		knapsack_backtrack2(items, curi+1, curW + items[curi], w)
	}
}

// 动态规划：二维数组保存操作结果
func knapsack3(items []int, w int) int {
	// 选择的结果记录
	states := make([][]bool, len(items))
	for i := 0; i < len(items); i++ {
		states[i] = append(states[i], make([]bool, w+1)...)
	}
	// 处理第一行数据
	states[0][0] = true
	if items[0] <= w {
		states[0][items[0]] = true
	}

	for i := 1; i < len(items); i++ {
		// 不选 i 的情况
		for j := 0; j < w+1; j++ {
			if states[i-1][j] {
				states[i][j] = true
			}
		}
		// 把 i 放入背包的情况
		for j := 0; j <= w - items[i]; j++ {
			if states[i-1][j] {
				states[i][j + items[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- { // 输出结果
		if states[len(items)-1][i] {
			return i
		}
	}
	return 0
}

// 动态规划：一维数组保存结果
func knapsack4(items []int, w int) int {
	// 选择的结果记录
	states := make([]bool, w+1)
	// 处理第一行数据
	states[0] = true
	if items[0] <= w {
		states[items[0]] = true
	}

	for i := 1; i < len(items); i++ {
		// 把第 i 个物品放入背包
		for j := w - items[i]; j >= 0; j-- {
			if states[j] {
				states[j + items[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[i] {
			return i
		}
	}

	return 0
}