package leetcode

import "sort"

func merge0056_0(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] <= intervals[j][0] {
			return true
		}
		return false
	})
	var merged [][]int
	for i := 0; i < len(intervals); i++ {
		// 左右端点
		L, R := intervals[i][0], intervals[i][1]
		// 如果当前区间的左端点在数组 merged 中最后一个区间的右端点之后，
		// 那么它们不会重合，我们可以直接将这个区间加入数组 merged 的末尾；
		if len(merged) == 0 || merged[len(merged)-1][1] < L {
			merged = append(merged, []int{L, R})
		} else {
			// 否则，它们重合，我们需要用当前区间的右端点更新数组 merged 中最后一个区间的右端点，
			// 将其置为二者的较大值。
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], R)
		}
	}

	return merged
}
