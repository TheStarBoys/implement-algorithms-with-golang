package leetcode

func isValidSudoku036_0(board [][]byte) bool {
	for x := range board {
		row := make(map[byte]int)
		for y := range board[x] {
			num := board[x][y]
			if num == '.' {
				continue
			}
			if _, ok := row[num]; ok {
				return false
			}
			row[num] = 1
		}
	}
	for y := range board[0] {
		col := make(map[byte]int)
		for x := range board {
			num := board[x][y]
			if num == '.' {
				continue
			}
			if _, ok := col[num]; ok {
				return false
			}
			col[num] = 1
		}
	}
	for i := 0; i < len(board); i= i+3 {
		for j := 0; j < len(board[0]); j = j+3 {
			if !findSmallTable(i, j, board) {
				return false
			}
		}
	}
	return true
}
func findSmallTable(x, y int, board [][]byte) bool {
	smallTable := make(map[byte]int)
	for i := x; i < x + 3; i++ {
		for j := y; j < y + 3; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}
			_, ok := smallTable[num]
			if ok {
				return false
			}
			smallTable[board[i][j]] = 1
		}
	}
	return true
}
