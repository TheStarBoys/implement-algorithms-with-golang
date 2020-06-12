package ByteDance

import (
	"bytes"
	"fmt"
)

// 给定类似www.toutiao.com的url，反转成com.toutiao.www的形式
// 要求时间O(N) 空间O(1)
// testcase1 ：www.51job.com.cn  --> cn.com.51job.www

//func main() {
//	url := []byte("www.51job.com.cn")
//	reverse(url)
//}

func reverse(url []byte) {
	helper(url, 0, len(url)-1) // reverse all
	//index1 := bytes.Index(url, []byte{'.'})
	index := -1
	preIndx := index
	for {
		preIndx++ // 0 --> 4 --> 12
		index = bytes.Index(url[preIndx:], []byte{'.'}) // 3 --> 7
		if index == -1 {
			break
		}
		index += preIndx // 3+0=3 --> 7+4=11
		helper(url, preIndx, index-1) // 0, 3 --> 4, 10
		preIndx = index // 3 --> 11
	}
	// preindx == 12
	helper(url, preIndx, len(url)-1)
	fmt.Println(string(url))
}
func helper(url []byte, l, r int) {
	for l < r {
		url[l], url[r] = url[r], url[l]
		l++
		r--
	}
}

