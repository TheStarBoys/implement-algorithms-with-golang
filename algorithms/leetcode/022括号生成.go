package leetcode

/*
// 效率太低
func generateParenthesis(n int) []string {
	max = n * 2
	ans = make([]string, 0)
	backtrace("", 0)
	return ans
}

var ans []string
var table = []byte{'(', ')'}
var max = 0
func backtrace(comb string, depth int) {
	if depth >= max {
		if isValid(comb) {
			ans = append(ans, comb)
		}
		return
	}
	for i := range table {
		backtrace(comb + string(table[i]), depth+1)
	}
}

func isValid(s string) bool {
	table := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	stack := make([]byte, 0)

	for i := range s {
		if v, ok := table[s[i]]; ok {
			if len(stack) == 0 { return false }
			if stack[len(stack) - 1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
*/
