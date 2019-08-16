package leetcode

import "bytes"

func reverseVowels(s string) string {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	b := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		i, j := bytes.IndexByte(vowels, b[l]), bytes.IndexByte(vowels, b[r])
		if i != -1 && j != -1 { // find
			b[l], b[r] = b[r], b[l]
			l++
			r--
		} else if i != -1{
			r--
		}else {
			l++
		}
	}
	return string(b)
}
