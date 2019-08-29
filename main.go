package main

import (
	"fmt"
)
func main() {
	n := 0
	fmt.Scanf("%d", &n)
	l, h := -90, 90
	count := 6
	for l + 1 < h && count > 0 {
		mid := (l + h) / 2
		if mid <= n {
			l = mid
			fmt.Print(1)
		}else {
			h = mid
			fmt.Print(0)
		}
		count--
	}
}