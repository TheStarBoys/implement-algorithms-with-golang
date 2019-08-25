package leetcode
// 双指针
func twoSum167_1(numbers []int, target int) []int {
	l, r := 0, len(numbers) -1
	sum := 0
	for l < r {
		if sum = numbers[l] + numbers[r]; sum == target {
			return []int{l+1, r+1}
		}else if sum < target {
			l++
		}else {
			r--
		}
	}
	return []int{}
}
// 二分查找，效率较低
func twoSum167_2(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{-1, -1}
	}
	for i := range numbers {
		rNum := target - numbers[i]
		r := binarySearch(i, len(numbers) - 1, rNum, numbers)
		if r != -1 {
			return []int{i + 1, r + 1}
		}
	}
	return []int{-1, -1}
}
// we know left, will find right
func binarySearch(l, h, target int, numbers []int) (r int) {
	for l + 1 < h {
		mid := l + (h - l) / 2
		if numbers[mid] == target {
			return mid
		}else if numbers[mid] > target {
			h = mid
		}else {
			l = mid
		}
	}
	if numbers[h] == target {
		return h
	}
	return -1
}