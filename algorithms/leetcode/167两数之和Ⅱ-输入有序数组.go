package leetcode

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) -1 // 双指针
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