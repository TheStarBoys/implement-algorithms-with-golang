package leetcode

func maxSubArray(nums []int) int {
	// 时间复杂度太高，不可取
	if len(nums) == 0 {
		return 0
	}
	// result set
	set := []int{}
	for i, v := range nums {
		count := 1
		for count <= len(nums) {
			sum := v
			for j := i+1; j < count; j++ {
				sum += nums[j]
			}
			count++
			set = append(set, sum)
		}
	}
	max := set[0]
	for _, v := range set {
		if v > max {
			max = v
		}
	}
	return max
}
//func maxSubArray(nums []int) int {
//
//}
func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}
