package leetcode

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := prices[0]
	earn := float64(0)
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy { // 当天prices比买入价格低，更新买入
			buy = prices[i]
		}else {
			earn = math.Max(float64(prices[i]-buy), earn)
		}
	}
	return int(earn)
}
