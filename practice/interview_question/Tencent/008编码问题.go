package Tencent

/*
假定一种编码的编码范围是a ~ y的25个字母，从1位到4位的编码，
如果我们把该编码按字典序排序，形成一个数组如下：
a, aa, aaa, aaaa, aaab, aaac, … …, b, ba, baa, baaa, baab, baac … …,
yyyw, yyyx, yyyy 其中a的Index为0，aa的Index为1，aaa的Index为2，
以此类推。 编写一个函数，输入是任意一个编码，输出这个编码对应的Index.
 */
/*
package main

import "fmt"

func main() {
	s := ""
	fmt.Scanf("%s", &s)
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b[i] = s[i]
	}
	decode(b)
}

func decode(s []byte) {
	sum, curSum := 0, 0
	n := len(s)
	for i := 0; i < 4; i++ {
		curSum *= 25
		if i < n {
			curSum += int(s[i]) - 'a'
		}
		sum += curSum
		if i < n - 1 {
			sum += 1
		}
	}
	fmt.Println(sum)
}

 */
