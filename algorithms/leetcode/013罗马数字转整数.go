package leetcode

/*
题目：
	给定一个罗马数字，将其转换成整数。
	输入确保在 1 到 3999 的范围内。
执行结果：
	用时：8 ms
	内存：3 MB
 */
// 假设输入的值是合法的
// 自己写的
func romanToInt(s string) int {
	// 用[]byte来保存每一个字符
	r, sum := []byte(s), 0
	m := map[byte]int{
		'I' : 1,
		'V' : 5,
		'X' : 10,
		'L' : 50,
		'C' : 100,
		'D' : 500,
		'M' : 1000,
	}
	for i := len(r) - 1; i >= 0; i-- {
		if i > 0 && m[r[i]] > m[r[i-1]] {	// 如果右侧的值比左侧的大
			sum += m[r[i]] - m[r[i-1]]// 等于原来的值 加上 右侧减左侧的值
			i--
			continue
		}
		switch r[i] {
		case byte('I') :
			sum += 1
		case byte('V') :
			sum += 5
		case byte('X') :
			sum += 10
		case byte('L') :
			sum += 50
		case byte('C') :
			sum += 100
		case byte('D') :
			sum += 500
		case byte('M') :
			sum += 1000
		}
	}
	return sum
}

func RomanToInt(s string) int {
	return romanToInt(s)
}