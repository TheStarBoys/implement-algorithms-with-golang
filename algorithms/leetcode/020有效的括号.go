package leetcode

/*
题目：
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
	有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	注意空字符串可被认为是有效字符串。
执行结果：
	用时：0 ms
	内存：2.1 MB
 */
func isValid(s string) bool {
	// [3]int 表示括号 "()" "{}" "[]"的闭合状况
	// 1 表示开， 0 表示闭合
	/*
	用栈来表示当前开着的括号，闭合后就出栈
	如果 开了之后，字符串完了，还没闭合， false
	如果 开了之后，中间有别的括号，没闭合外层括号就闭合了，false
	 */
	if s == "" {
		return true
	}
	bracket := [3]int{}
	stack := []string{}
	bs := []byte(s)
	if b := bs[0]; b == byte(')') || b == byte('}') || b == byte(']') {
		return false
	}
	for _, b := range bs {
		switch b {
		case byte('(') :
			bracket[0]++
			stack = append(stack, "()")
		case byte(')') :
			bracket[0]--
			if checkStack("()", stack) {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		case byte('{') :
			bracket[1]++
			stack = append(stack, "{}")
		case byte('}') :
			bracket[1]--
			if checkStack("{}", stack) {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		case byte('[') :
			bracket[2]++
			stack = append(stack, "[]")
		case byte(']') :
			bracket[2]--
			if checkStack("[]", stack) {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}
	for _, v := range bracket {
		if v != 0 {
			return false
		}
	}
	return true
}
// 校验当前闭合的括号跟栈顶的元素是否相同
func checkStack(bracket string, stack []string) bool{
	if len(stack) == 0 {
		return false
	}
	if bracket != stack[len(stack) - 1] {
		return false
	}
	return true
}

func IsValid(s string) bool {
	return isValid(s)
}