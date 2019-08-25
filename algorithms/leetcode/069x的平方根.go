package leetcode

// 通用模板
func mySqrt0(x int) int {
	if x <= 1 {
		return x
	}
	l, h := 0, x
	for l + 1 < h {
		mid := l + (h-l)/2
		sqrt := x / mid
		if mid == sqrt {
			return mid
		}else if mid > sqrt {
			h = mid
		}else {
			l = mid
		}
	}
	return l
}

func mySqrt1(x int) int {
	if x <= 1 {
		return x
	}
	l, h := 1, x
	for l <= h {
		mid := l + (h - l)/2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		}else if sqrt > mid {
			l = mid + 1
		}else {
			h = mid - 1
		}
	}
	return h
}