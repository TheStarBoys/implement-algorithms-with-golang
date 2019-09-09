package leetcode
// 解题思路：
// 先遍历nums计算出所有的和，复制给rightSum
// 在遍历nums，rightSum减去当前下标的值==当前的rightSum
// 如果左边的和跟右边的和不相等，leftSum += nums[i]
func pivotIndex724_0(nums []int) int {
	leftSum, rightSum := 0, 0
	for _, v := range nums {
		rightSum += v
	}
	for i := range nums {
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
