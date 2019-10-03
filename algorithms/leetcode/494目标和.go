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
func findTargetSumWays(nums []int, S int) int {
	return dfs(nums, 0, 0, S)
}

func dfs(nums []int, indx int, count, S int) int {
	if indx >= len(nums) {
		if S == 0 {
			count++
		}
		return count
	}
	count = dfs(nums, indx+1, count, S-nums[indx])
	return dfs(nums, indx+1, count, S+nums[indx])
}