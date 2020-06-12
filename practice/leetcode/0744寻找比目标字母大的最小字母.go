package leetcode

// 效率似乎跟普通解法无太大差异
func nextGreatestLetter(letters []byte, target byte) byte {
	l, h := 0, len(letters) - 1
	for l <= h {
		mid := l + (h - l) / 2
		if letters[mid] <= target {
			l = mid + 1
		}else {
			h = mid - 1
		}
	}
	if l > len(letters) - 1 {
		return letters[0]
	}
	return letters[l]
}

// 最普通的解法
func nextGreatestLetter0(letters []byte, target byte) byte {
	for i := range letters {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}
