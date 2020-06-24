package leetcode

// dfs
func findCircleNum0547_0(M [][]int) int {
	var ans int
	visited := map[int]bool{} // 顶点是否被访问过
	for i := 0; i < len(M); i++ { // 遍历顶点
		if visited[i] == false {
			ans++
			dfs0547_0(i, visited, M)
		}
	}
	return ans
}

func dfs0547_0(i int, visited map[int]bool, M [][]int) {
	visited[i] = true
	for j := 0; j < len(M); j++ {
		if !visited[j] && M[i][j] == 1 {
			dfs0547_0(j, visited, M)
		}
	}
}
