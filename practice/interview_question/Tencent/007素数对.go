package Tencent

/*
给定一个正整数，编写程序计算有多少对质数的和等于输入的这个正整数，并输出结果。输入值小于1000。
如，输入为10, 程序应该输出结果为2。（共有两对质数的和为10,分别为(5,5),(3,7)）
 */

/*

import "fmt"
func main() {
	n := 0
	fmt.Scanf("%d", &n)
	nums := make([]int, 0)
	Flag:
	for i := 2; i < n; i++ {
		for j := i - 1; j > 1; j-- {
			if i % j == 0 {
				continue Flag
			}
		}
		nums = append(nums, i)
	}
	count := 0
	used := make(map[int]bool)
	Flag2:
	for i := range nums {
		if _, ok := used[nums[i]]; ok {
			continue
		}
		l, h := 0, len(nums) - 1
		target := n - nums[i]
		for l + 1 < h {
			mid := l + (h - l) / 2
			if nums[mid] == target {
				count++
				used[nums[i]] = true
				used[nums[mid]] = true
				continue Flag2
			}else if nums[mid] > target {
				h = mid
			}else {
				l = mid
			}
		}
		if nums[l] == target {
			used[nums[i]] = true
			used[nums[l]] = true
			count++
			continue
		}else if nums[h] == target {
			used[nums[i]] = true
			used[nums[h]] = true
			count++
			continue
		}
	}
	fmt.Println(count)
}


 */
