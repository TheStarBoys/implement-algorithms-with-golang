package leetcode

// 滑动窗口
func checkInclusion0567_0(s1 string, s2 string) bool {
	hash := [26]int{}
	for i := 0; i < len(s1); i++ {
		hash[s1[i] - 'a']++
	}
	for l, r := 0, 0; r < len(s2); r++ {
		hash[s2[r] - 'a']--
		for hash[s2[r] - 'a'] < 0 {
			hash[s2[l] - 'a']++
			l++
		}
		if r - l + 1 == len(s1) {
			return true
		}
	}
	return false
}
