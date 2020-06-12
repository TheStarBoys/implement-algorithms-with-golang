package leetcode


// 这道题有个隐藏条件，题目给定条件一定能跳到最后，否则这样解会死循环
func jump045_0(nums []int) int {
	count := 0
	// 找到当前最大的实际跳跃长度
	for i := 0; i < len(nums)-1; { // 此处不能自增
		maxLength := 0
		indx := 0
		if i + nums[i] >= len(nums)-1 { // 考虑[3,2,1]
			return count+1
		}
		for j := i+1; j <= i + nums[i] && j < len(nums); j++ {
			// j 是将要跳到的下标
			// j - i + nums[j] 是当前实际跳跃长度
			// fmt.Println("i =", i, "j =", j, "len =", j - i + nums[j])
			if j - i + nums[j] > maxLength {
				maxLength = j - i + nums[j]
				indx = j
			}
		}
		// 跳到局部最优解处
		i = indx
		count++ // 考虑[0]
		// fmt.Printf("max = %d, indx=%d\n", maxLength, indx)
	}
	return count
}

