package Tencent

import "math"

/*
小Q的父母要出差N天，走之前给小Q留下了M块巧克力。
小Q决定每天吃的巧克力数量不少于前一天吃的一半，但是他又不想在父母回来之前的某一天没有巧克力吃，
请问他第一天最多能吃多少块巧克力
 */
/*
package main

import (
    "fmt"
    "math"
)

func main() {
    n, m := 0, 0
    fmt.Scanf("%d %d", &n, &m)
    fmt.Println(eatChocolate(n, m))
}

 */
func eatChocolate(n, m int) int {
	l, h := 0, m
	for l <= h {
		mid := l + (h-l)/2
		sum := first(mid, n)
		if sum == m {
			return mid
		}else if sum > m {
			h = mid - 1
		}else {
			l = mid + 1
		}
	}
	return h
}
// 第一天的c巧克力, 一共n天
func first(c, n int) int {
	pre := c
	for i := 1; i < n; i++ {
		pre = int(math.Ceil(float64(pre) / 2))
		c += pre
	}
	return c
}