package Tencent
import (
	"fmt"
	"strings"
)

// 问题描述见 pic对应题号
func main() {
	t := 0
	fmt.Scan(&t)
	for i := 0; i < t; i ++ {
		n := 0
		fmt.Scan(&n)
		s := ""
		fmt.Scan(&s)
		if isQQ(s) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
func isQQ(s string) bool {
	indx :=strings.Index(s, "8")
	if indx == -1 {
		return false
	}
	return len(s) - indx >= 11
}
