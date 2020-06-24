package leetcode

import "bytes"

// 迭代：超时
func kthGrammar0779_0(N int, K int) int {
	bts := []byte{'0'}
	for i := 0; i < N; i++ {
		tmp := make([][]byte, len(bts))
		for j := range bts {
			if bts[j] == '0' {
				tmp[j] = []byte("01")
			} else {
				tmp[j] = []byte("10")
			}
		}
		bts = bytes.Join(tmp, []byte{})
	}

	return int(bts[K-1] - '0')
}

// 递归
func kthGrammar0779_1(N int, K int) int {
	// 如果我们已知 dp(n, k), 我们可以用 O(1) 的时间求出 dp(n+1, 2k) 和 dp(n+1, 2k+1)
	// 反过来说 求 dp(n, k), 如果 k 是奇数并且 dp(n-1, k/2) 为 0，则 dp(n, k) 为 1
	// 如果 k 是偶数 并且 dp(n-1, k/2) 为1，则 dp(n, k) 为 1
	var recursive func(n, k int) int
	recursive = func(n, k int) int {
		if n == 0 && k == 0 { return 0 }
		prev := recursive(n-1, k/2)
		if prev == 1 && k & 1 == 0 ||
			prev == 0 && k & 1 == 1 {
			return 1
		}
		return 0
	}
	return recursive(N-1, K-1)
}