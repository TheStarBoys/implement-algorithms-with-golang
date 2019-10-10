package leetcode
// 待优化
func findTargetSumWays494_0(nums []int, S int) int {
	DFS494_0(nums, 0, 0, S)
	return result
}
var result int
func DFS494_0(nums []int, index, cS, S int) {
	if cS == S && len(nums) == index {
		result++
		return
	}
	if index >= len(nums) {
		return
	}
	// +
	DFS494_0(nums, index + 1, cS + nums[index], S)
	// -
	DFS494_0(nums, index + 1, cS - nums[index], S)
}

// 放弃全局变量，因为OJ用全局变量可能带来问题，减少Cs的传参
func findTargetSumWays494_1(nums []int, S int) int {
	return dfs494_1(nums, 0, 0, S)
}

func dfs494_1(nums []int, indx int, count, S int) int {
	if indx >= len(nums) {
		if S == 0 {
			count++
		}
		return count
	}
	count = dfs494_1(nums, indx+1, count, S-nums[indx])
	return dfs494_1(nums, indx+1, count, S+nums[indx])
}

// 思路：queue里存放遍历数组后所得到结果的每一种可能
// 空间复杂度O(2^N) 时间复杂度O(2^n - 2) = O(2^n)
func findTargetSumWays494_2(nums []int, S int) int {
	if len(nums) == 0 { return -1 }
	queue := []int{nums[0], -nums[0]}
	p := 1
	for len(queue) > 0 && p < len(nums) {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

			queue = append(queue, cur+nums[p])
			queue = append(queue, cur-nums[p])
		}
		p++
	}
	count := 0
	for _, v := range queue {
		if v == S {
			count++
		}
	}
	return count
}

// 时间复杂度比第494_2少一半
func findTargetSumWays494_3(nums []int, S int) int {
	if len(nums) == 0 { return -1 }
	queue := []int{nums[0], -nums[0]}
	count := 0
	if len(nums) == 1 {
		if S == queue[0] {
			count++
		}
		if S == queue[1] {
			count++
		}
		return count
	}
	p := 1
	for len(queue) > 0 && p < len(nums) {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			if p == len(nums) - 1 {
				if cur + nums[p] == S {
					count++
				}
				if cur - nums[p] == S {
					count++
				}
				continue
			}
			queue = append(queue, cur+nums[p])
			queue = append(queue, cur-nums[p])
		}
		p++
	}
	return count
}