package leetcode

// dfs解法，隐式调用系统栈，必须用哈希集来确定节点是否被访问过，否则会堆栈溢出
func floodFill733_0(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	visited := make(map[string]bool)
	image[sr][sc] = newColor
	visited[k([]int{sr, sc})] = true
	dfs733_0(image, sr, sc, oldColor, newColor, visited)
	return image
}
//var dx = []int{0, 0, -1, 1}
//var dy = []int{-1, 1, 0, 0}
func dfs733_0(image [][]int, i, j, oldColor, newColor int, visited map[string]bool) {
	for z := range dx {
		temp_i := i + dx[z]
		temp_j := j + dy[z]
		if 0 <= temp_i && temp_i < len(image) && 0 <= temp_j && temp_j < len(image[0]) &&
			!visited[k([]int{temp_i, temp_j})] && image[temp_i][temp_j] == oldColor {
			image[temp_i][temp_j] = newColor
			visited[k([]int{temp_i, temp_j})] = true
			dfs733_0(image, temp_i, temp_j, oldColor, newColor, visited)
		}
	}
}
//func k(list []int) string {
//	return fmt.Sprintf("%q", list)
//}

// bfs解法
func floodFill733_1(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	visited := make(map[string]bool)
	bfs733_1(image, sr, sc, oldColor, newColor, visited)
	return image
}
//var dx = []int{0, 0, -1, 1}
//var dy = []int{-1, 1, 0, 0}
func bfs733_1(image [][]int, i, j, oldColor, newColor int, visited map[string]bool) {
	queue := make([][]int, 0)
	queue = append(queue, []int{i, j})
	visited[k([]int{i, j})] = true
	image[i][j] = newColor

	for len(queue) != 0 {
		length := len(queue)
		for z := 0; z < length; z++ {
			curIndex := queue[0]
			i, j = curIndex[0], curIndex[1]
			queue = queue[1:]

			for h := range dx {
				temp_i := i + dx[h]
				temp_j := j + dy[h]
				if 0 <= temp_i && temp_i < len(image) && 0 <= temp_j && temp_j < len(image[0]) &&
					!visited[k([]int{temp_i, temp_j})] && image[temp_i][temp_j] == oldColor {
					image[temp_i][temp_j] = newColor
					visited[k([]int{temp_i, temp_j})] = true
					queue = append(queue, []int{temp_i, temp_j})
				}
			}
		}
	}

}
//func k(list []int) string {
//	return fmt.Sprintf("%q", list)
//}

// dfs解法，显示调用栈
func floodFill733_2(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	visited := make(map[string]bool)
	dfs733_2(image, sr, sc, oldColor, newColor, visited)
	return image
}
//var dx = []int{0, 0, -1, 1}
//var dy = []int{-1, 1, 0, 0}
func dfs733_2(image [][]int, i, j, oldColor, newColor int, visited map[string]bool) {
	stack := make([][]int, 0)
	stack = append(stack, []int{i, j})
	visited[k([]int{i, j})] = true
	image[i][j] = newColor
	for len(stack) != 0 {
		cur := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		i, j = cur[0], cur[1]
		for z := range dx {
			temp_i := i + dx[z]
			temp_j := j + dy[z]
			if 0 <= temp_i && temp_i < len(image) && 0 <= temp_j && temp_j < len(image[0]) &&
				!visited[k([]int{temp_i, temp_j})] && image[temp_i][temp_j] == oldColor {
				visited[k([]int{temp_i, temp_j})] = true
				stack = append(stack, []int{temp_i, temp_j})
				image[temp_i][temp_j] = newColor
			}
		}
	}
}
//func k(list []int) string {
//	return fmt.Sprintf("%q", list)
//}