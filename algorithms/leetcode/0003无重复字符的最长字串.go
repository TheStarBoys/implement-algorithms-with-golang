package leetcode

import "math"

// 自己的解法, 效率太低
func lengthOfLongestSubstring003_0(s string) int {
	table := make(map[byte]int)
	maxLength := 0
	curLength := 0
	for i := 0; i < len(s); i++ {
		if index, ok := table[s[i]]; ok {
			if maxLength < curLength {
				maxLength = curLength
			}
			curLength = 0
			i = index + 1
			table = make(map[byte]int)
		}
		curLength++
		table[s[i]] = i
	}
	if curLength != 0 && maxLength < curLength {
		return curLength
	}
	return maxLength
}
// sliding window
func lengthOfLongestSubstring003_1(s string) int {
	ans, i, j := 0, 0, 0
	n := len(s)
	set := make(map[byte]bool)

	for i < n && j < n {
		if !set[s[j]] {
			set[s[j]] = true
			j++
			ans = int(math.Max(float64(ans), float64(j - i)))
		} else {
			delete(set, s[i])
			i++
		}
	}
	return ans
}

// sliding window 优化
func lengthOfLongestSubstring003_2(s string) int {
	n, ans := len(s), 0
	table := make(map[byte]int)
	for i, j := 0, 0; j < n; j ++ {
		if index, ok := table[s[j]]; ok {
			i = int(math.Max(float64(index), float64(i)))
		}
		ans = int(math.Max(float64(ans), float64(j - i + 1)))
		table[s[j]] = j + 1
	}
	return ans
}