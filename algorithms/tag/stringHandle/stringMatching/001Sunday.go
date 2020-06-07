package stringMatching

import "strings"

// sunday算法，字符串模式匹配
func SundaySearch(str, sub string) int{
	cur := 0 // 指向当前str和sub对齐的位置
	stri, subi := 0, 0 // 双指针，分别指向str, sub
	Flag :
	for stri < len(str) && subi < len(sub) && cur+len(sub) <= len(str){
		if str[stri] != sub[subi] {
			// 匹配失败，关注src中参加匹配的最末位字符的下一个字符
			i := 0
			for i = len(sub)-1; i >= 0; i-- {
				if cur+len(sub) < len(str) && sub[i] == str[cur+len(sub)] {
					// 末位的下一个字符在sub中
					cur = cur + (len(sub) - i)
					stri = cur
					subi = 0
					continue Flag
				}
			}
			if i < 0 {
				// 末位的下一个字符不在sub中
				cur = cur + len(sub) + 1
				stri = cur
				subi = 0
				continue
			}
		}
		stri++
		subi++
	}
	if subi >= len(sub) {
		return cur
	}
	return -1
}

// 去掉所有空格
func MyTrimSpace(s string) string{
	return strings.ReplaceAll(s, " ", "")
}
