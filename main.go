package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1}, 3))
}

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