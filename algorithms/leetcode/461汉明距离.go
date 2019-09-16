package leetcode

// 将汉明距离通过异或运算，转换成计算位1的个数
// 计算位1的个数的方案，请跳转至 191题参考
func hammingDistance461_0(x int, y int) int {
	x ^= y
	bits := 0
	mask := 1
	for i := 0; i < 32; i++ {
		if (x & mask) != 0 {
			bits++
		}
		mask <<= 1
	}
	return bits
}

func hammingDistance461_1(x int, y int) int {
	x ^= y
	sum := 0
	for x != 0 {
		sum++
		x &= x-1
	}
	return sum
}