package leetcode

// n = b^x  ==> n = b * b * b * ... * b
// 如果n是3的幂，那么n不断的除以3，最终结果应该等于1
func isPowerOfThree326_0(n int) bool {
	if n < 1 { return false }
	for n % 3 == 0 {
		n /= 3
	}
	return n == 1
}
// 整数限制
// n的最大值为 1 << 31 - 1,  3 ^ i = n  --> log3(n) = i
// 3的幂的最大值=3 ^ log3(n) = 3^19 = 1162261467
// 如果n是3的幂，那么n一定是3^0, 3^1 ... 3^19中的一个
// 所以判断3^19 % n == 0就知道是不是3的幂了
func isPowerOfThree326_1(n int) bool {
	return n > 0 && 1162261467 % n == 0
}