package leetcode

// dfs，隐式调用系统栈
func canVisitAllRooms841_0(rooms [][]int) bool {
	visited := make(map[int]bool)
	dfs841_0(rooms, 0, visited)

	for i := range rooms {
		if !visited[i] {
			return false
		}
	}
	return true
}
func dfs841_0(rooms [][]int, room int, visited map[int]bool) {
	visited[room] = true
	for _, k := range rooms[room] {
		if !visited[k] {
			visited[k] = true
			dfs841_0(rooms, k, visited)
		}
	}
}
