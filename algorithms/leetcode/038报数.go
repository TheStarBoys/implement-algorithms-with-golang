package leetcode

import "strconv"

/*
从左往右找到与该元素相同的个数，记为，[个数 元素值]
直到所有的元素都找一遍
 */

func countAndSay(n int) string {
	// 快慢指针
	cur, fast := 0, 0
	serial, temp := "1", ""
	if n == 1 {
		return serial
	}
	for i := 1; i < n; i++ {
		for fast < len(serial) {
			if serial[cur] != serial[fast] {
				temp += strconv.Itoa(fast-cur) + string(serial[cur])
				cur = fast
			}
			fast++
			if fast == len(serial) {
				temp += strconv.Itoa(fast-cur) + string(serial[cur])
			}
		}
		serial = temp
		cur, fast, temp = 0, 0, ""
	}
	return serial
}
// 递归
func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}
	pre := countAndSay(n-1)
	count := 1
	now := "" // 现在的结果
	for i := 0; i < len(pre); i++ {
		if i == len(pre) -1 {
			now += strconv.Itoa(count) + string(pre[i])
			count = 1
			continue
		}
		if pre[i] == pre[i+1] {
			count++
			continue
		}else {
			now += strconv.Itoa(count) + string(pre[i])
			count = 1
		}
	}
	return now
}
func CountAndSay(n int) string {
	return countAndSay(n)
}