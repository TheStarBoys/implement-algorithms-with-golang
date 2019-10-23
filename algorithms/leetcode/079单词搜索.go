package leetcode

import "fmt"

func exist079_0(board [][]byte, word string) bool {
	if word == "" {
		return false
	}
	for i := range board {
		for j := range board[0] {
			if word[0] == board[i][j] {
				visited := make(map[string]bool)
				visited[helper079_0([]int{i, j})] = true
				if dfs079_0(visited, board, word, i, j, 1) { return true }
			}
		}
	}
	return false
}

//var dx = []int{0, 0, -1, 1}
//var dy = []int{-1, 1, 0, 0}
func dfs079_0(visited map[string]bool, board [][]byte, word string, i, j, next int) bool {
	if next > len(word)-1 {
		return true
	}
	for x := range dx {
		temp_i := dx[x] + i
		temp_j := dy[x] + j
		if 0 <= temp_i && temp_i < len(board) &&
			0 <= temp_j && temp_j < len(board[0]) &&
			!visited[helper079_0([]int{temp_i, temp_j})] &&
			board[temp_i][temp_j] == word[next] {
			if next == len(word)-1 {
				return true
			}
			visited[helper079_0([]int{temp_i, temp_j})] = true
			if dfs079_0(visited, board, word, temp_i, temp_j, next+1) { return true }
			// 不能缺少，这是回溯的过程，进行了状态的重置。
			// 如果没有此回溯，示例：
			// exist([][]byte{
			//		[]byte{'a', 'b', 'c', 'e'},
			//		[]byte{'s', 'f', 'e', 's'},
			//		[]byte{'a', 'd', 'e', 'e'},
			//	}, "abceseeefs")
			// 的结果为false，其实应该为true
			delete(visited, helper079_0([]int{temp_i, temp_j}))
		}
	}
	return false
}

func helper079_0(list []int) string {
	return fmt.Sprintf("%q", list)
}
