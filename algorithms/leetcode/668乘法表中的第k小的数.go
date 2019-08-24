package leetcode

import "sort"
/*
如果小于mid的数超过了k，说明mid大了
 */
func findKthNumber(m int, n int, k int) int {
	l, h := 1, m*n
	for l + 1 < h {
		mid := l + (h-l)/2
		if count(mid, m, n, k) {
			h = mid
		}else {
			l = mid
		}
	}
	if count(l, m, n, k) {
		return l
	}
	return h
}
/*
i 是当前行数
tmp := mid / i
tmp 是当前行小于mid的数量
 */
func count(mid, m, n, k int) bool {
	res := 0
	for i := 1; i <= m; i++ {
		tmp := mid / i
		if tmp == 0 {
			break
		}
		if tmp < n {
			res += tmp
		}else {
			res += n
		}
	}
	return res >= k
}

// 超出内存限制
func findKthNumber0(m int, n int, k int) int {
	if k == 1 {
		return 1
	}
	table := make([][]int, m)
	for i := range table {
		table[i] = make([]int, n)
	}
	// j == 0 || i == 0  --> return i+1 || j+1
	// f(i, j) = f(i, 0) * f(0, j)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			helper668(i, j, table)
		}
	}
	h := []int{}
	for _, v := range table {
		h = append(h, v...)
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i] < h[j]
	})
	return h[k-1]
}
func helper668(i, j int, table [][]int) int {
	if j == 0 {
		table[i][j] = i+1
	}else if i == 0 {
		table[i][j] = j+1
	}
	if table[i][j] != 0 {
		return table[i][j]
	}
	table[i][j] = helper668(i, 0, table) * helper668(0, j, table)
	return table[i][j]
}