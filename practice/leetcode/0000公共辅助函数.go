package leetcode

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 快速乘
// 如果将 B 展开为二进制，其下标 i 对最后结果做的贡献为 A * (1 << i)，即 A << i
// 遍历 B 二进制展开后的每一位进行贡献的累加，最后的结果就是答案
// 例如：A: 2 B: 3
// 2 --> 0010 3 --> 0011 6 --> 0110
// i == 0 贡献为 0010 --> 2
// i == 1 贡献为 0100 --> 4
// i == 2 无贡献
// i == 3 无贡献
// 最终结果 2 + 4 = 6

// 注意：该实现 B 不能为负数
func quickMulti(A, B int) int {
	ans := 0
	for ; B > 0; B >>= 1 {
		if B & 1 > 0 {
			ans += A
		}
		A <<= 1
	}

	return ans
}

// 支持 B 为负数
func quickMulti2(A, B int) int {
	ans := 0
	// 用于将数 n 正负性转换
	reverse := func(n *int) bool {
		*n = -*n
		return true
	}

	// 如果满足 B < 0，将 B 转为 -B，flag标识 B 已经被转换
	flag := B < 0 && reverse(&B)
	// 快速乘迭代
	for ; B > 0; B >>= 1 {
		if B & 1 > 0 {
			ans += A
		}
		A <<= 1
	}

	// 如果 B 被转换过，将最终答案正负性进行转换
	_ = flag && reverse(&ans)
	return ans
}

// 交换切片 i, j 下标的值
func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}