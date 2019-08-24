package leetcode

// 二分查找行和列
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	length := len(matrix[0])
	l, h := 0, len(matrix)-1
	i := 0
	for l + 1 < h {
		mid := l + (h-l)/2
		if matrix[mid][length-1] == target {
			return true
		}else if matrix[mid][length-1] > target {
			h = mid
		}else {
			l = mid
		}
	}
	if matrix[l][length-1] >= target {
		i = l
	}else if  matrix[h][length-1] >= target {
		i = h
	}
	if matrix[i][length-1] == target {
		return true
	} else if matrix[i][length-1] > target {
		l, h := 0, length-1
		for l+1 < h {
			mid := l + (h-l)/2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				h = mid
			} else {
				l = mid
			}
		}
		if target == matrix[i][l] {
			return true
		} else if target == matrix[i][h] {
			return true
		}
	}
	return false
}
// 二分查找列，遍历查找行
func searchMatrix0(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0 {
			return false
		}
		if matrix[i][len(matrix[i])-1] == target {
			return true
		} else if matrix[i][len(matrix[i])-1] > target {
			l, h := 0, len(matrix[i])-1
			for l+1 < h {
				mid := l + (h-l)/2
				if matrix[i][mid] == target {
					return true
				} else if matrix[i][mid] > target {
					h = mid
				} else {
					l = mid
				}
			}
			if target == matrix[i][l] {
				return true
			} else if target == matrix[i][h] {
				return true
			}
		}
	}
	return false
}
