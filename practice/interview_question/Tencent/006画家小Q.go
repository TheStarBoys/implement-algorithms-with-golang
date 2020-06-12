package Tencent

import "fmt"

/*
解题思路：
X = 0, Y = 1, B = 2, G = 3
遍历到!0的元素时，判断是1还是2，如果是3就跳过。这一次遍历，是去除所有非3的元素
1：向右下角遍历，确定这一笔画的长度；并在这些元素上-1，将画的笔数+1
2：向左下角遍历，确定这一笔画的长度；并在这些元素上-1，将画的笔数+1
第二次遍历：如果元素=3，判断左下遍历还是右下遍历，元素-3，笔数+2，如果都不满足，笔数+2

一般情形：
1,0,0,2
0,1,3,0
0,2,1,1
2,0,0,1

特殊情形：
3,0,0,0
0,0,3,0
0,3,0,3
0,0,0,0
*/
// board = [n][m]int
/*
/*
测试用例：
28 12
BYGGYYXYXGGG
BYBBXYBBBYGX
YYXBXYBGGXBB
XYGYGXBBBGBB
YBXGGXYYBXBY
GXGXXBBXYBBG
XBBXBGXXXYXY
YXGYBBGYXXGG
GXGYGGXXYBXY
GXBBXGXBXYGB
XBBGYYYBXGXG
YGGBYGXBYXBX
YYXXXGYXBGXB
XYXYXYBGGXGB
XGYYYGYXYYGB
YBYGGYGXGBGX
YXBXYGGXYXYG
GGYGXXBYBXGG
GBBBYGGBYXGX
XBYYXGBBGXGB
YBYXBBBYGXBY
YYBGXYYBXBXY
YBXBYYBGXGBY
XYGGGXXYGGYG
BBBBYBGYXYBG
GXBXXGYGYXYG
XBYYGGBBBBXG
GXYBXGYXYGGY
*/
/*
func main() {
	//board := [][]int{
	//	{1, 0, 0, 2},
	//	{0, 1, 3, 0},
	//	{0, 2, 1, 1},
	//	{2, 0, 0, 1},
	//}
	//board := [][]int{
	//	{3, 0, 0, 0},
	//	{0, 0, 3, 0},
	//	{0, 3, 0, 3},
	//	{0, 0, 0, 0},
	//}
	n, m := 0, 0
	fmt.Scanf("%d %d", &n, &m)
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		str := ""
		fmt.Scanf("%s\n", &str)
		for _, v := range str {
			//X = 0, Y = 1, B = 2, G = 3
			if v == 'X' {
				board[i] = append(board[i], 0)
			}else if v == 'Y' {
				board[i] = append(board[i], 1)
			}else if v == 'B' {
				board[i] = append(board[i], 2)
			}else if v == 'G' {
				board[i] = append(board[i], 3)
			}
		}
	}
	printBoard(board)
	fmt.Println(painter(board))
}
*/
// 通过率60%
func painter(board [][]int) int {
	count := 0
	for !only3(board) {
		for i := range board {
			for j := range board[i] {
				if board[i][j] == 1 {
					drawUp(i, j, 1, board)
					drawDown(i, j, 1, board)
					fmt.Println("count = ", count)
					printBoard(board)
					count++
				}
				if board[i][j] == 2 {
					drawUp(i, j, 2, board)
					drawDown(i, j, 2, board)
					fmt.Println("count = ", count)
					printBoard(board)
					count++
				}
			}
		}
	}
	// 第二次遍历
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 3 {
				drawDown(i, j, 3, board)
				fmt.Println("count = ", count)
				printBoard(board)
				count += 2
			}
		}
	}
	return count
}
func drawDown(i, j, pen int, board [][]int) {
	for i < len(board) && j < len(board[i]) && j >= 0 && (board[i][j] == pen || board[i][j] == 3 ) {
		board[i][j] -= pen
		if pen == 1 {
			j++
		}else if pen == 2 {
			j--
		}else if pen == 3 {
			if i+1 < len(board) {
				if j-1 >= 0 && board[i+1][j-1] == 3 {
					j--
				}else if j+1 < len(board[i]) && board[i+1][j+1] == 3 {
					j++
				}
			}
		}
		i++
	}
	return
}
func drawUp(i, j, pen int, board [][]int) {
	if pen == 1 {
		j--
	}else if pen == 2 {
		j++
	}
	i--
	for i >= 0 && j < len(board[i]) && j >= 0 && (board[i][j] == pen || board[i][j] == 3 ) {
		board[i][j] -= pen
		if pen == 1 {
			j--
		}else if pen == 2 {
			j++
		}
		i--
	}
	return
}
func only3(board [][]int) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 1 || board[i][j] == 2 {
				return false
			}
		}
	}
	return true
}
func printBoard(board [][]int) {
	for i := range board {
		fmt.Println(board[i])
	}
}