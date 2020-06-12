package leetcode

import "sort"

// 哈希表
// 用一个数组来记录已出现的数字，置为1
// 时间O(n) 空间O(n)
func missingNumber268_0(nums []int) int {
	table := make([]int, len(nums) + 1)
	for _, v := range nums {
		table[v] = 1
	}
	for i, v := range table {
		if v == 0 {
			return i
		}
	}
	return -1
}
// 数学
// 解题思路：
// 例子：[3,0,1],其中n=3, 0~n的累加和=6, 而数组的累加和3+0+1=4
// 6 - 4 = 2, 而2恰好是我们要求的缺失的数字
// 总结式子：total - sum = target
// 其中total是0~n的累加和，sum是数组所有元素和，target就是返回值
// 0~n的累加和可以用数学公式直接算出来total = n*(n+1)/2

// 时间O(n) 空间O(1)
func missingNumber268_1(nums []int) int {
	n := len(nums)
	total := n*(n+1)/2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return total - sum
}
// 异或
// 未缺失的数字，在异或的算式中成对出现，这两个相同的数字异或后，等于0
// 缺失的数字，没有配对的数字跟它异或，所以最终的结果，就是缺失的数字

// 时间O(n) 空间O(1)
func missingNumber268_2(nums []int) int {
	miss := len(nums)
	for i := range nums {
		miss ^= i ^ nums[i]
	}
	return miss
}

// 排序
// 时间复杂度最高，不推荐
func missingNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] != n {
		return n
	}
	if nums[0] != 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return -1
}