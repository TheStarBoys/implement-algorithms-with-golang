package stringMatching

import "math"

// 时间复杂度：O(n)
func RKSearch(str, sub string) int {
	if len(str) < len(sub) {
		return -1
	}
	var (
		m = len(sub)
		subCode = hashCode3(sub)
		l = 0
		// code 是滑动窗口中子串的哈希值
		code = hashCode3(str[:m])
	)
	// 利用滑动窗口，可以用O(1)的时间从 h[i-1] 递推得到 h[i]，i 为子串起始下标
	for r := m; ; r++ {
		if code == subCode {
			// 如果哈希算法是允许有哈希冲突的，此处还需要额外比较子串跟模式串是否相等
			//if str[l : r] == sub {
			//	return l
			//}
			return l
		}
		if r >= len(str) {
			break
		}
		//code -= int(str[l] - 'a' + 1) * int(math.Pow(26, float64(m-1)))
		code -= int(str[l] - 'a' + 1) * pow(m-1)
		l++
		code *= 26
		code += int(str[r] - 'a' + 1)
	}

	return -1
}

// 基于 进制转换的思想，a ~ z 一共26个字母，把它当成一个26进制，"不存在哈希冲突"(其实a从0开始导致会有一定问题)
// 这种方式其值很容易就超出了整型的范围
// 有一定问题，字符串 "abc" 跟 "bc" 结果其实是一样的，都是28
func hashCode1(str string) int {
	sum := 0
	i := 0
	for j := len(str) - 1; j >= 0; j-- {
		sum += int(str[j] - 'a') * int(math.Pow(26, float64(i)))
		i++
	}

	return sum
}

// 基于第一种方式，将 a 记为 1 而不是0
func hashCode2(str string) int {
	sum := 0
	i := 0
	for j := len(str) - 1; j >= 0; j-- {
		sum += int(str[j] - 'a' + 1) * int(math.Pow(26, float64(i)))
		i++
	}

	return sum
}

// 查表法优化 math.Pow()
func hashCode3(str string) int {
	sum := 0
	i := 0
	for j := len(str) - 1; j >= 0; j-- {
		sum += int(str[j] - 'a' + 1) * pow(i)
		i++
	}

	return sum
}

// pow 可以用查表法优化
var powTable =  []int{
	0: 1,
	1: 26,
	2: 676,
	3: 17576,
	4: 456976,
	5: 11881376,
	6: 308915776,
	7: 8031810176,
	8: 208827064576,
	9: 5429503678976,
	10: 141167095653376,
	11: 3670344486987776,
	12: 95428956661682176,
	13: 2481152873203736576,
}
// pow returns 26 ^ x
func pow(x int) int {
	if x < 0 {
		panic("x less than zero")
	}
	if x > 13 {
		panic("exceed int range")
	}

	return powTable[x]
}