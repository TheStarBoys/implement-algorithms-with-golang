package ByteDance

/*
错误解答
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n, m, q := 0, 0, 0
	fmt.Scan(&n, &m, &q)
	cubes := make([]byte, n)
	s := ""
	tmp := ""
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		s += tmp
	}
	strings.ReplaceAll(s, " ", "")
	cubes = []byte(s)
	cubesTmp := make([]byte, n)
	l, r := 0, 0
	for i := 0; i < q; i++ { // q次游戏
		fmt.Scan(&l, &r)
		copy(cubesTmp, cubes)
		fmt.Println(getScore(cubesTmp, l, r))
		l, r = 0, 0
	}
}
func getScore(cubes []byte, l, r int) int {
	direction := 1	// 0 向左， 1 向右
	visit := 0 // 如果刚访问过方向符， visit = 1
	score := 0
	for i := l - 1; i < r && i >= l - 1; {
		if '0' < cubes[i] && cubes[i] <= '9' { // == 0 暂时没做处理
			if visit == 1 {
				visit = 0
			}
			sco, _ := strconv.Atoi(string(cubes[i]))
			score += sco
			sco--
			cubes[i] = byte(sco + '0')
		} else if cubes[i] == '<' {
			if visit == 1 {
				cubes[i] = '0'
			}
			visit = 1
			direction = 0
		} else if cubes[i] == '>' {
			if visit == 1 {
				cubes[i] = '0'
			}
			visit = 1
			direction = 1
		}
		if direction == 0 {
			i--
		} else {
			i++
		}
	}
	return score
}


 */
