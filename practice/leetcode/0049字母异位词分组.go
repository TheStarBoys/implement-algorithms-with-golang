package leetcode

import "sort"

// 时间复杂度胜97%
func groupAnagrams049_0(strs []string) [][]string {
	table := make(map[string][]string)
	for _, str := range strs {
		key := getKey049_0(str)
		table[key] = append(table[key], str)
	}
	ans := [][]string{}
	for _, v := range table {
		ans = append(ans, v)
	}
	return ans
}
func getKey049_0(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

// 略微优化时间
func groupAnagrams049_1(strs []string) [][]string {
	table := make(map[string]int) // 给字母异位词分一个数组空间, int对应的数组下标
	var indx int
	ans := [][]string{}
	for _, str := range strs {
		key := getKey049_1(str)
		if i, ok := table[key]; ok {
			ans[i] = append(ans[i], str)
			continue
		}
		table[key] = indx // 分配下标
		ans = append(ans, []string{}) // 并分配空间
		ans[indx] = append(ans[indx], str)
		indx++
	}
	return ans
}
func getKey049_1(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

