package byte字节跳动


/*

错误解答
package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}
	fmt.Println(getMax(heights))
}

func getMax(heights []int) int {
	table := make(map[int]int)	// height - > count
	for i := len(heights) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if heights[j] >= heights[i] {
				table[heights[j]]++
				break
			}
		}
	}
	table1 := make(map[int]int) // count -> height
	maxCount := 0
	for k, v := range table {
		table1[v] = k
		if maxCount < v {
			maxCount = v
		}
	}
	return table1[maxCount]
}
 */
