package leetcode

// 遍历查找 时间O(n^2)
func findDuplicate287_0(nums []int) int {
	//n := len(nums) - 1
	for i := range nums {
		for j := i+1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[j]
			}
		}
	}
	return -1
}
// 利用哈希表，时间O(n), 空间O(n)
func findDuplicate287_1(nums []int) int {
	table := make([]int, len(nums))
	for _, v := range nums {
		table[v]++
	}
	for i, v := range table {
		if v > 1 {
			return i
		}
	}
	return -1
}
// 哈希表优化，一次遍历 时间O(n), 空间O(n)
func findDuplicate287_2(nums []int) int {
	table := make([]int, len(nums))
	for _, v := range nums {
		if table[v] != 0 {
			return v
		}
		table[v]++
	}
	return -1
}

// 鸽子洞原理， 将这个数组转换为链表的思维来看
// 下标是链表的Val，值是链表的Next指针
// 如果存在重复元素，那么必定会形成环
// 此时用快慢指针就可以找到是否有环，并根据此找到环的入口，入口即是重复的数

// 此解法对应的图，我放在自己OneNote笔记 --> 算法 --> LeetCode287里
func findDuplicate287_3(nums []int) int {
	// Find the intersection of the two runners.
	tortoise := nums[0] // 龟 --> 慢指针,从一开始就指向next
	hare := nums[0] // 兔 --> 快指针,从一开始就指向next
	flag := 1
	for flag == 1 || tortoise != hare {
		flag--
		tortoise = nums[tortoise] // 慢指针移动一次
		hare = nums[nums[hare]] // 快指针移动两次
	}
	// Find the "entrance" to the cycle.
	hare = tortoise
	tortoise = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}
	return tortoise
}

// 二进制
func findDuplicate287_4(nums []int) int {
	n := len(nums)
	ans := 0
	bit_max := uint(31)
	// 找到最高有效位
	for ((n-1) >> bit_max) == 0 {
		bit_max--
	}

	for bit := uint(0); bit <= bit_max; bit++ {
		// 考虑到第 i 位，我们记 nums[] 数组中二进制展开后第 i 位为 1 的数有 x 个
		// 数字 [1,n] 这 n 个数二进制展开后第 i 位为 1 的数有 y 个
		// 那么重复的数第 i 位为 1 当且仅当 x > y
		x, y := 0, 0
		for i := 0; i < n; i++ {
			if (nums[i] & (1 << bit)) > 0 {
				x++
			}
			if i >= 1 && (i & (1 << bit)) > 0 {
				y++
			}
		}
		// 按位还原
		if x > y {
			ans |= 1 << bit
		}
	}
	return ans
}