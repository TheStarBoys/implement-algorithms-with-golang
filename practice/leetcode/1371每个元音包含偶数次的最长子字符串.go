package leetcode

// 暴力解，超时
func findTheLongestSubstring1371_0(s string) int {
	if s == "" {
		return 0
	}
	hash := [5]int{}
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			increaseCount1371_0(s[j], &hash)
			if isEven1371_0(hash) {
				res = max(res, i-j+1)
			}
		}
		hash = [5]int{}
	}
	return res
}

func increaseCount1371_0(b byte, hash *[5]int) {
	switch b {
	case 'a':
		hash[0]++
	case 'e':
		hash[1]++
	case 'i':
		hash[2]++
	case 'o':
		hash[3]++
	case 'u':
		hash[4]++
	}
}

func isEven1371_0(hash [5]int) bool {
	for _, n := range hash {
		if n % 2 != 0 {
			return false
		}
	}
	return true
}


