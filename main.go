package main

import (
	"fmt"
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
	fmt.Println(findKthNumber(2, 3, 6))
	//fmt.Println(count(4, 2, 3, 6))
}
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

func printRow(r []int) {
	for _, v := range r {
		fmt.Print(v, " ")
	}
	fmt.Println()
}