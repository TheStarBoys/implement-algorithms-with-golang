package backtrack

type RExp struct {
	matched bool
	pattern string // 正则表达式
	plen int // 正则表达式长度
}

func NewRegularExpresstion(pattern string) RExp {
	return RExp{
		pattern: pattern,
		plen: len(pattern),
	}
}

func (r *RExp) Match(text string) bool {
	r.matched = false
	r.rmatch(0, 0, text, len(text))

	return r.matched
}

// tlen 文本串长度
func (r *RExp) rmatch(ti, pj int, text string, tlen int) {
	if r.matched { return } // 如果已经匹配，则不继续递归
	if pj == r.plen { // 正则表达式到结尾了
		if ti == tlen { // 此时文本串也结尾，说明匹配
			r.matched = true
		}
		return
	}
	switch r.pattern[pj] {
	case '*':
		// 匹配任意个字符
		// 注意这里的循环条件为小于等于
		// 等于是为了解决 ti 到末尾，而表达式还剩一个 '*' 的情况
		for i := ti; i <= tlen; i++ {
			r.rmatch(i, pj+1, text, tlen)
		}
	case '?':
		// 匹配零个或多个字符
		r.rmatch(ti, pj+1, text, tlen)
		r.rmatch(ti+1, pj+1, text, tlen)
	default:
		// 普通字符，需要匹配才继续递归
		if ti < tlen && text[ti] == r.pattern[pj] {
			r.rmatch(ti+1, pj+1, text, tlen)
		}
	}
}