package main

import (
	"fmt"
	"math"
)

/*
5
1 2 3 3 5
3
1 2 1
2 4 5
3 5 3
 */
func main() {
	nums := []int{1, -1}
	fmt.Println(maxSlidingWindow(nums, 1))
}
func maxSlidingWindow(nums []int, k int) []int {
	maxs := []int{}
	for i := range nums {
		if i + k > len(nums) {
			break
		}
		maxs = append(maxs, slidWindow(i, i+k, nums))
	}
	return maxs
}
// 0000 0001
func slidWindow(i, j int, nums []int) int {
	max := float64(-1 << 63)
	for x := i; x < j; x++ {
		max = math.Max(max, float64(nums[x]))
	}
	return int(max)
}