package ByteDance

/*
5 2 3
3 1 2 3
0
2 2 3
1 2
1 3
*/
/*

https://www.nowcoder.com/question/next?pid=8537209&qid=141021&tid=26496289

func main() {
	n, m, c := 0, 0, 0
	fmt.Scanf("%d %d %d", &n, &m, &c)
	cuan := make([][]int, n)
	for i := 0; i < n; i++ {
		colorNum := 0
		fmt.Scanf("%d", &colorNum)
		for j := 0; j < colorNum; j++ {
			color := 0
			fmt.Scanf("%d", &color)
			cuan[i] = append(cuan[i], color)
		}
	}
	count := 0
	errColor := make([]int, c+1)
	for i := 0; i < n; i++ {
		for j := 0; j < len(cuan[i]); j++ {
			if 1 == errColor[cuan[i][j]] {
				continue
			}
			if hasSameColor(cuan[i][j], i, m, cuan) {
				errColor[cuan[i][j]] = 1
			}
		}
	}
	for _, v := range errColor {
		if v == 1 {
			count++
		}
	}
	fmt.Println(count)
}

 */
func hasSameColor(color, num, m int, cuan [][]int) bool {
	if m+num > len(cuan) {
		for i := num + 1; i < len(cuan); i++ {
			for j := range cuan[i] {
				if cuan[i][j] == color {
					return true
				}
			}
		}
		for i := 0; i < m+num-len(cuan); i++ {
			for j := range cuan[i] {
				if cuan[i][j] == color {
					return true
				}
			}
		}
		return false
	}
	for i := num + 1; i < m+num; i++ {
		for j := range cuan[i] {
			if cuan[i][j] == color {
				return true
			}
		}
	}
	return false
}
