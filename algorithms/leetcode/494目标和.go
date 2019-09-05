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
