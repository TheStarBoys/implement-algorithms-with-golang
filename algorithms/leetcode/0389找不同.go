package leetcode

import "sort"

// 排序遍历法
func findTheDifference389_0(s string, t string) byte {
	// acd --> abcd
	s1, t1 := []byte(s), []byte(t)
	helper389_0(s1)
	helper389_0(t1)
	for i := range s1 {
		if s1[i] != t1[i] {
			return t1[i]
		}
	}
	return t1[len(t1)-1]
}

func helper389_0(s []byte) {
	sort.Slice(s, func(i,j int) bool {
		return s[i] < s[j]
	})
}


// 哈希表法
func findTheDifference389_1(s string, t string) byte {
	// acd --> abcd
	table := make(map[byte]int)
	for i := range s {
		table[s[i]]++
	}
	for i := range t {
		if n := table[t[i]]; n > 0 {
			table[t[i]]--
		} else {
			return t[i]
		}
	}
	return '*'
}


// 异或法
func findTheDifference389_2(s string, t string) byte {
	// acd --> abcd
	b := t[len(t)-1]
	for i := range s {
		b = ((b-'a') ^ (s[i]-'a') ^ (t[i]-'a')) + 'a'
	}
	return b
}


// 相加求差法
func findTheDifference389_3(s string, t string) byte {
	// acd --> abcd
	sCount := 0
	tCount := int(t[len(t)-1])
	for i := range s {
		sCount += int(s[i])
		tCount += int(t[i])
	}
	return byte(tCount-sCount)
}

