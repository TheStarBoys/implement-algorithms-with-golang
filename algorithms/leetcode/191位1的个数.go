package leetcode

func hammingWeight191_0(num uint32) int {
	bits := 0
	mask := uint32(1)
	for i := 0; i < 32; i++ {
		if (num & mask) != uint32(0) {
			bits++
		}
		mask <<= 1
	}
	return bits
}
// 最优方案
func hammingWeight191_1(num uint32) int {
	sum := 0
	for num != uint32(0) {
		sum++
		num &= num - 1
	}
	return sum
}