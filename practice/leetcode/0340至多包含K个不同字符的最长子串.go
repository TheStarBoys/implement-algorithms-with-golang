package leetcode

import "math"


// Sliding Window
func lengthOfLongestSubstringKDistinct0340_1(s string, k int) int {
	if s == "" {
		return 0
	}
	// 范围不是小写字母
	hash := make(map[byte]int)
	res, count, l := 0, 0, 0
	for r := 0; r < len(s); r++ {
		hash[s[r]]++
		// 新增的不同字符
		if hash[s[r]] == 1 {
			count++
		}
		for count > k {
			hash[s[l]]--
			if hash[s[l]] == 0 {
				count--
			}
			l++
		}
		if count <= k {
			res = int(math.Max(float64(res), float64(r - l + 1)))
		}
	}

	return res
}
