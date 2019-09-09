package byte字节跳动
// 时间复杂度过高， 没通过

//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	n, D := 0, 0
//	fmt.Scan(&n, &D)
//	street := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&street[i])
//	}
//	fmt.Println(getWays(street, D))
//}
//func getWays(street []int, D int) int {
//	if len(street) <= 2 {
//		return 0
//	}
//	for i := range street {
//		dfs(street, D, i, street[i], 1)
//	}
//	return result
//}
//var result int
//func dfs(street []int, D, i, oldSite, deep int) {
//	if street[i] - oldSite > D {
//		return
//	}
//	if deep >= 3 {
//		result++
//		return
//	}
//	for j := i+1; j < len(street); j++ {
//		dfs(street, D, j, oldSite, deep + 1)
//	}
//}
