package offer

// 滑动窗口
func findContinuousSequence(target int) [][]int {
	var ans [][]int
	l := 1
	count := 0
	for r := 1; r < target; r++ {
		count += r
		for count > target {
			count -= l
			l++
		}
		if count == target {
			tmp := []int{}
			for i := l; i <= r; i++ {
				tmp = append(tmp, i)
			}
			ans = append(ans, tmp)
		}
	}
	return ans
}
