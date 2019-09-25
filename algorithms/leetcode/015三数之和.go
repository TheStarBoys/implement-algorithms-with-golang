package leetcode

import (
	"fmt"
	"sort"
)
// 解法1：暴力遍历，超时
func threeSum015_0(nums []int) [][]int {
	ans := [][]int{}
	table := make(map[string]bool)
	for i := range nums {
		for j := range nums {
			for k := range nums {
				if i == j || i == k || j == k {
					continue
				}
				if nums[i] + nums[j] +nums[k] == 0 {
					tmp := []int{nums[i], nums[j], nums[k]}
					sort.Ints(tmp)
					if !table[h015(tmp)] {
						table[h015(tmp)] = true
						ans = append(ans, tmp)
					}

				}
			}
		}
	}
	return ans
}
func h015(list []int) string {
	return fmt.Sprintf("%q", list)
}

// 解法2：时间O(n^2)
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := range nums {
		if nums[i] > 0 { return ans } // 如果当前的元素都大于0了，那么后面就没有相加后能等于0的元素了
		if i > 0 && nums[i] == nums[i-1] { continue } // 去重
		L, R := i+1, len(nums) - 1 // 注意L不是从0开始的
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] { L++ } // 去重
				for L < R && nums[R] == nums[R-1] { R-- } // 去重
				L++
				R--
			} else if sum > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return ans
}