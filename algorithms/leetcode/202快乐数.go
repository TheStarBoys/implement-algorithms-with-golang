package leetcode

// 解题思路：如果会出现无限循环的情况，那么必定会出现
// 当前的n与之前已经计算过的n相同，继而不断的执行相同的步骤，无法退出循环
// 故而需要维护一个哈希集，来保证n的值是没访问过的
func isHappy202_0(n int) bool {
	set := make(map[string]bool)
	for {
		temp := 0
		pair := []int{}
		for n != 0 {
			num := n % 10
			pair = append(pair, num)
			temp += num * num
			n /= 10
		}
		if set[k(pair)] { return false }
		set[k(pair)] = true
		if temp == 1 { return true }
		n = temp
	}
}
//func k(list []int) string {
//	return fmt.Sprintf("%q", list)
//}

// 优化，实际上不需要一个名为pair的[]int变量来存储所有的位，直接存储n的值即可
func isHappy202_1(n int) bool {
	set := make(map[int]bool)
	set[n] = true
	for {
		temp := 0
		for n != 0 {
			num := n % 10
			temp += num * num
			n /= 10
		}
		if temp == 1 { return true }
		if set[temp] { return false }
		set[temp] = true
		n = temp
	}
}