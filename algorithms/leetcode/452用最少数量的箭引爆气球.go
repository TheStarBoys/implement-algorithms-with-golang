package leetcode

import "sort"
// 解题思路 :
// 刨除没有气球，使用0根箭的情况
// 将所有气球坐标按从小到大的顺序排列
// 有大于等于一个气球的时候，至少需要一根箭，所以count := 1
// end用来存储当前重叠的区间的最小右区间
// 如果i元素的左区间小于等于当前重叠区间的右区间 points[i][0] <= end (说明区间重叠)
// 此时，如果points[i][1] < end, 则更新最小重叠的右区间end
// 如果区间不重叠, end = points[i][1], 并且需要射出的箭矢+1
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	count := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= end { // 重叠
			if points[i][1] < end {
				end = points[i][1]
			}
		} else {
			end = points[i][1]
			count++
		}
	}
	return count
}
