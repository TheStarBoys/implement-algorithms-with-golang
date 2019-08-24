package leetcode

// 二分查找， 理解思想，但不建议这样做
func searchMatrixV2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	numOfFind := 0
	if len(matrix) > len(matrix[0]) {
		numOfFind = len(matrix[0])
	}else {
		numOfFind = len(matrix)
	}
	for i := 0; i < numOfFind; i++ {
		upDown := find(matrix, i, target, true)
		leftRight := find(matrix, i, target, false)
		if upDown || leftRight {
			return true
		}
	}
	return false
}
func find(matrix [][]int, i, target int, isFindRow bool) bool {
	l, h := i, 0
	if isFindRow {
		h = len(matrix) - 1
	}else {
		h = len(matrix[0]) - 1
	}
	for l + 1 < h {
		mid := l + (h-l)/2
		if isFindRow {
			if matrix[mid][i] < target {
				l = mid
			}else if matrix[mid][i] > target {
				h = mid
			}else {
				return true
			}
		} else {
			if matrix[i][mid] < target {
				l = mid
			}else if matrix[i][mid] > target {
				h = mid
			}else {
				return true
			}
		}
	}
	if isFindRow {
		if matrix[l][i] == target || matrix[h][i] == target {
			return true
		}
	}else {
		if matrix[i][l] == target || matrix[i][h] == target {
			return true
		}
	}
	return false
}
// 贪心思想
func searchMatrixV2_1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix[0]) - 1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}else if matrix[i][j] > target {
			j--
		}else {
			i++
		}
	}
	return false
}
