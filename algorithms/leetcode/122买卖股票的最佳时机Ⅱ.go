package leetcode

func maxProfit122_0(prices []int) int {
	if len(prices) == 0 { return 0 }
	buy := 0
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] <= prices[buy] {
			buy = i
		} else {
			sum += prices[i] - prices[buy]
			buy = i
		}
	}
	return sum
}
// 优化0的方案
func maxProfit122_1(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
