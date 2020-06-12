package leetcode

// 解法1：空间最优，操作局部优化（双指针）
func moveZeroes283_0(nums []int)  {
	i := 0
	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++
		}
	}
	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
}
// 解法2：最优解，是解法1的细微扩展
func moveZeroes283_1(nums []int)  {
	last := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[last], nums[cur] = nums[cur], nums[last]
			last++
		}
	}
}