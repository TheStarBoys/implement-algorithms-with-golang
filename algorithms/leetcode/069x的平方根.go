package leetcode

func mySqrt(x int) int {
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