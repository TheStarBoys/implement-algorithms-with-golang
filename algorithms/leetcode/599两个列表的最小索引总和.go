package leetcode

import "math"

// 时间O(2N+M)
func findRestaurant599_0(list1 []string, list2 []string) []string {
	table1 := make(map[string]int) // restaurant --> index
	table2 := make(map[string]int)
	tableIndx := make(map[int][]string)
	minIndex := math.MaxInt64
	for i, v := range list1 {
		table1[v] = i
	}
	for i, v := range list2 {
		table2[v] = i
	}
	for k, v := range table1 {
		if v2, ok := table2[k]; ok { // 如果有相同的餐厅
			sum := v + v2
			if minIndex > sum { // 找到最小索引和
				minIndex = sum
			}
			tableIndx[sum] = append(tableIndx[sum], k)
		}
	}
	return tableIndx[minIndex]
}
// 空间优化：不用table2了，时间优化：两次循环遍历即可
// 时间O(N+M)
func findRestaurant599_1(list1 []string, list2 []string) []string {
	table1 := make(map[string]int) // restaurant --> index
	tableIndx := make(map[int][]string)
	minIndex := math.MaxInt64
	for i, v := range list1 {
		table1[v] = i
	}
	for i, v := range list2 {
		if indx, ok := table1[v]; ok { // 如果有相同的餐厅
			sum := indx + i
			if minIndex > sum {
				minIndex = sum
			}
			tableIndx[sum] = append(tableIndx[sum], v)
		}
	}
	return tableIndx[minIndex]
}