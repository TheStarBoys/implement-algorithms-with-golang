package leetcode

func sortedArrayToBST0108(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	var root TreeNode
	root.Val = nums[mid]
	root.Left = sortedArrayToBST0108(nums[:mid])
	root.Right = sortedArrayToBST0108(nums[mid+1:])
	return &root
}
