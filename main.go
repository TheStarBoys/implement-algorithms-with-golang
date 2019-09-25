package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"ad", "da", "cd"}))
}

func groupAnagrams(strs []string) [][]string {
	table := make(map[string][]string)
	for _, str := range strs {
		key := getKey(str)
		table[key] = append(table[key], str)
	}
	ans := [][]string{}
	for _, v := range table {
		ans = append(ans, v)
	}
	return ans
}
func getKey(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}