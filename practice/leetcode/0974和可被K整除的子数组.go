package leetcode

// 暴力解：超时
func subarraysDivByK0974_0(A []int, K int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += A[j]
			if sum % K == 0 {
				count++
			}
		}
	}

	return count
}

// 哈希表 + 逐一统计

// 时间复杂度：O(N)，其中 N 是数组 A 的长度。
// 我们只需要从前往后遍历一次数组，在遍历数组的过程中，
// 维护哈希表的各个操作均为 O(1)，因此总时间复杂度为 O(N)

// 空间复杂度：O(min(N,K))，即哈希表需要的空间。
// 当 N <= K 时，最多有 N 个前缀和，因此哈希表中最多有 N+1 个键值对；
// 当 N >= K 时，最多有 K 个不同的余数，因此哈希表中最多有 K 个键值对。
// 也就是说，哈希表需要的空间取决于 N 和 K 中的较小值。
func subarraysDivByK0974_1(A []int, K int) int {
	record := map[int]int{0:1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum % K + K) % K
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}