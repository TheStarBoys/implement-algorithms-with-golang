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
			if !findSmallTable036_0(i, j, board) {
				return false
			}
		}
	}
	return true
}
func findSmallTable036_0(x, y int, board [][]byte) bool {
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

// 哈希表， 一次迭代
func isValidSudoku036_1(board [][]byte) bool {
	row := make([]map[int]bool, 9)
	col := make([]map[int]bool, 9)
	boxes := make([]map[int]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make(map[int]bool)
		col[i] = make(map[int]bool)
		boxes[i] = make(map[int]bool)
	}
	for i := range board {
		for j, v := range board[i] {
			if v < '0' || v > '9' {
				continue
			}
			num := int(v - '0')
			if row[i][num] {
				return false
			}
			row[i][num] = true
			if col[j][num] {
				return false
			}
			col[j][num] = true
			box_index := (i / 3) * 3 + j / 3
			if boxes[box_index][num] {
				return false
			}
			boxes[box_index][num] = true
		}
	}
	return true
}