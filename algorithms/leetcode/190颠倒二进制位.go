package leetcode


func reverseBits190_0(num uint32) uint32 {
	mask := uint32(1) // 用mask来跟num的每一位相与，得到res的最后一位的值
	res := num & mask
	for i := 1; i < 32; i++ {
		mask <<= 1
		res <<= 1
		if num & mask != 0 {
			res += 1
		}
	}
	return res
}
// 思想与上面的方法类似，但重复利用了num，减少了不必要的mask的内存
func reverseBits190_1(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res <<= 1
		res += num&1
		num >>= 1
	}
	return res
}