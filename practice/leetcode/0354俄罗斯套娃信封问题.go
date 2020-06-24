package leetcode

import "sort"

// 排序后贪心：解答错误，局部最优解无法推导出全局最优解，只能是接近最优
func maxEnvelopes0354_0(envelopes [][]int) int {
	if len(envelopes) == 0 { return 0 }
	// sort ==> [[2,3],[5,4],[6,4],[6,7]]
	sort.Slice(envelopes, func(i, j int) bool{
		if envelopes[i][0] < envelopes[j][0] {
			return true
		}
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] <= envelopes[j][1]
		}
		return false
	})
	// fmt.Printf("sorted: %v\n", envelopes)
	prevIndx, ans := 0, 1 // 注意：ans初值是1
	for i := 1; i < len(envelopes); i++ {
		// fmt.Printf("prev: %v, curr: %v, ans: %d\n", envelopes[prevIndx], envelopes[i], ans)
		if envelopes[prevIndx][0] < envelopes[i][0] &&
			envelopes[prevIndx][1] < envelopes[i][1] {
			ans++
			prevIndx = i
		}
	}

	return ans
}

// 回溯超时
func maxEnvelopes0354_1(envelopes [][]int) int {
	if len(envelopes) == 0 { return 0 }
	var (
		ans int
		backtrack func(curr, count int)
	)
	backtrack = func(curr, count int) {
		if count > ans {
			ans = count
		}
		for i := 0; i < len(envelopes); i++ {
			if i != curr && envelopes[curr][0] < envelopes[i][0] &&
				envelopes[curr][1] < envelopes[i][1] {
				backtrack(i, count+1)
			}
		}
	}
	for i := 0; i < len(envelopes); i++ {
		backtrack(i, 1)
	}

	return ans
}

// 回溯剪枝依旧超时
func maxEnvelopes0354_2(envelopes [][]int) int {
	if len(envelopes) == 0 { return 0 }
	var (
		ans int
		backtrack func(curr, count int)
		mem = make([]int, len(envelopes))
	)
	backtrack = func(curr, count int) {
		if mem[curr] > 0 {
			if mem[curr] > count {
				return
			}
		}
		mem[curr] = count
		if count > ans {
			ans = count
		}
		for i := 0; i < len(envelopes); i++ {
			if i != curr && envelopes[curr][0] < envelopes[i][0] &&
				envelopes[curr][1] < envelopes[i][1] {
				backtrack(i, count+1)
			}
		}
	}
	for i := 0; i < len(envelopes); i++ {
		backtrack(i, 1)
	}

	return ans
}

// 先排序再回溯：79 / 85 个通过测试用例
func maxEnvelopes0354_3(envelopes [][]int) int {
	if len(envelopes) == 0 { return 0 }
	sort.Slice(envelopes, func(i, j int) bool{
		if envelopes[i][0] < envelopes[j][0] {
			return true
		}
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] <= envelopes[j][1]
		}
		return false
	})
	var (
		ans int
		backtrack func(curr, count int)
		mem = make([]int, len(envelopes))
	)
	backtrack = func(curr, count int) {
		if mem[curr] > 0 {
			if mem[curr] > count {
				return
			}
		}
		mem[curr] = count
		if count > ans {
			ans = count
		}
		for i := curr+1; i < len(envelopes); i++ {
			if envelopes[curr][0] < envelopes[i][0] &&
				envelopes[curr][1] < envelopes[i][1] {
				backtrack(i, count+1)
			}
		}
	}
	for i := 0; i < len(envelopes); i++ {
		backtrack(i, 1)
	}

	return ans
}

