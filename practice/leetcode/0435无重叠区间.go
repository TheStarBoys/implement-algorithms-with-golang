package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end { // 无重叠区间
			end = intervals[i][1]
		} else {
			if intervals[i][1] < end { // 区间重叠，且当前区间更优
				end = intervals[i][1]
			}
			count++
		}
	}
	return count
}


func EraseOverlapIntervals(intervals [][]int) int {
	return eraseOverlapIntervals(intervals)
}