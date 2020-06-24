package leetcode

// 时间复杂度：O(N * K) K是连续序列长度的平均值
func longestConsecutive0128_0(nums []int) int {
	table := map[int]bool{}
	ans := 0
	for _, v := range nums {
		table[v] = true
	}
	for _, v := range nums {
		count := 1
		for ; table[v+1]; v, count = v+1, count+1 {}
		ans = max(ans, count)
	}

	return ans
}

// 时间复杂度：O(N) 因为在满足是连续序列的第一个数时才会进入内层循环
func longestConsecutive0128_1(nums []int) int {
	table := map[int]bool{}
	ans := 0
	for _, v := range nums {
		table[v] = true
	}
	for _, v := range nums {
		if !table[v-1] {
			count := 1
			for ; table[v+1]; v, count = v+1, count+1 {}
			ans = max(ans, count)
		}
	}

	return ans
}