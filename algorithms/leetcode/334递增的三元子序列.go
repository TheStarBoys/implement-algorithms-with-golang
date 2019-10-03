package leetcode

import "math"

// 这是一个错误版本，由于审错了题
//func increasingTriplet(nums []int) bool {
//	if len(nums) < 3 { return false }
//	l, pre, r := 0, 0, 1
//	for r < len(nums) && r - l < 3 {
//		if nums[r] <= nums[pre] {
//			l, pre, r = r, r, r+1
//			continue
//		}
//		r++
//		pre++
//	}
//	if r - l >= 3 { return true }
//	return false
//}
// todo 暂时没有看懂
func increasingTriplet334_0(nums []int) bool {
	n1 := math.MaxInt32
	n2 := n1
	n3 := n1
	for _, v := range nums {
		if v < n1 {
			if n2 == math.MaxInt32 {
				n1 = v
			} else if v < n3 {
				n3 = v
			} else {
				n1, n2, n3 = n3, v, math.MaxInt32
			}
		} else if v > n1 {
			if v < n2 {
				n2 = v
			} else if v > n2 { return true }
		}
	}
	return false
}











