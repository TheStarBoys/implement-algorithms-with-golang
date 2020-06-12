package leetcode

import (
	"math"
	"strconv"
)

// 转成数字来做，到后面num会大到超过int范围
func removeKdigits0402_1(num string, k int) string {
	a, _ := strconv.Atoi(num)
	i := removeNumKdigits0402_1(a, k)

	return strconv.Itoa(i)
}

// 单调栈
func removeKdigits0402_2(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	stack := []byte{}

	// 形成一个单调栈
	for i := range num {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	// 如果删除的数字m比k少，那么从尾部再删除k-m个数字
	// for i := 0; i < k; i++ {
	//     stack = stack[:len(stack)-1]
	// }
	stack = stack[:len(stack) - k]
	// 去掉前导零
	i := 0
	for ; i < len(stack) && stack[i] == '0'; i++ {}
	stack = stack[i:]

	// 可能从序列中删除所有的数字
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}


// 类似题，不过不是给的字符串num，而是一个非负整数 a

// 用回溯来做，会超时
func removeNumKdigits0402_1(a int, k int) int {
	if k <= 0 {
		return a
	}
	e := 10
	mod := 1
	m := math.MaxInt64
	for a >= mod {
		// 1234 4
		// 1234 / 10 = 123
		// 1234 % 1 = 0
		// 123 = 123 * 1 + 0

		// 1234 3
		// 1234 / 100 = 12
		// 1234 % 10 = 4
		// 124 = 12 * 10 + 4

		// 1234 2
		// 1234 / 1000 = 1
		// 1234 % 100 = 34
		// 134 = 1 * 100 + 34

		// 1234 1
		// 1234 / 10000 = 0
		// 1234 % 1000 = 234
		// 234 = 0 * 1000 + 234

		// 总结得到:
		// a / e * mod + a % mod
		// 这个式子可以去掉a中相应位置的数字
		m = min(m, removeNumKdigits0402_1(a / e * mod + a % mod, k-1))
		mod = e
		e *= 10
	}
	return m
}

// 贪心来做
func removeNumKdigits0402_2(a int, k int) int {
	if k <= 0 {
		return a
	}
	// 直接移除掉了所有的数字
	if int(math.Pow(10, float64(k))) > a {
		return 0
	}
	// 每次做的选择，都将获得局部的最小值
	// 做完K次选择后，则获得到了全局的最小值
	minVal := a
	for i := 0; i < k; i++ {
		tmp := minVal
		e := 10
		mod := 1
		for tmp >= mod {
			// 获得局部最优解
			minVal = min(minVal, tmp / e * mod + tmp % mod)
			mod = e
			e *= 10
		}
	}

	return minVal
}