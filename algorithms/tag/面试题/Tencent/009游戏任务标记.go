package Tencent

// 32个unsigned int， int == 32位， 32个 * 32位 == 1024
// 题意：将数从int -- > [32]int里面存放, nums := [32]int{}
// 由于1 ~ 1024 对应计算机的位是 0 ~ 1023，所以刨除临界条件后
// n1-- n2--
// 给定一个数n1, 例如：100
// 所在下标 index 1 := n1 / 32
// 所在下标偏移量 offset1 := n1 % 32
// 最终数 n1 表示为 ：nums[index1] = 1 << offset1
// 把对应位置 1 说明任务已完成
// n2 = 100
// index2 := n2 / 32
// offset2 := n2 % 32
// nums[index2] == 1 << offset2 如果相等，说明该任务已完成

/*
package main
import "fmt"
func main() {
	n1, n2 := 0, 0
	fmt.Scanf("%d %d", &n1, &n2)
	fmt.Println(isComplete(n1, n2))
}
func isComplete(n1, n2 int) int {
	if n1 > 1024 || n2 > 1024 || n1 < 1 || n2 < 1 {
		return -1
	}
	n1--
	n2--
	nums := [32]uint{}
	index1 := n1 / 32
	offset1 := n1 % 32
	index2 := n2 / 32
	offset2 := n2 % 32
	nums[index1] = 1 << uint(offset1)
	if nums[index2] == 1 << uint(offset2) {
		return 1
	}
	return 0
}

 */

