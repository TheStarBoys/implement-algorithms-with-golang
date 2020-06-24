package offer

func firstUniqChar(s string) byte {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	for i := range s {
		if m[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
