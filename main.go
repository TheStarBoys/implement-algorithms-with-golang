package main

import (
	"fmt"
)

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1, 3},{3, 0, 1}, {2}, {0}}))
}
func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	dfs(rooms, 0, visited)

	for i := range rooms {
		if !visited[i] {
			return false
		}
	}
	return true
}
func dfs(rooms [][]int, room int, visited map[int]bool) {
	visited[room] = true
	for _, k := range rooms[room] {
		if !visited[k] {
			visited[k] = true
			dfs(rooms, k, visited)
		}
	}
}



