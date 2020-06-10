package ac_automation

// AC自动机
type ACNode struct {
	children map[rune]*ACNode
	isWord bool
	length int // 当 trie.Search() == true时，记录模式串长度
	fail *ACNode // 失败指针
}

func New() *ACNode {
	return &ACNode{
		children: map[rune]*ACNode{},
	}
}

func (ac *ACNode) Insert(word []rune) {
	trie := ac
	for _, r := range word {
		next, ok := trie.children[r]
		if !ok {
			tmp := New()
			next = tmp
			trie.children[r] = next
		}
		trie = next
	}
	trie.isWord = true
	trie.length = len(word)
}

func (ac *ACNode) Search(word []rune) bool {
	trie := ac
	for _, r := range word {
		next, ok := trie.children[r]
		if !ok {
			return false
		}
		trie = next
	}

	return trie.isWord || len(trie.children) == 0
}

func (root *ACNode) BuildFailurePointer() {
	var queue []*ACNode
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			p := queue[0]
			queue = queue[1:]
			for data, pc := range p.children {
				if p == root {
					pc.fail = root
				} else {
					q := p.fail
					for q != nil {
						qc, ok := q.children[data]
						if ok {
							pc.fail = qc
							break
						}
						q = q.fail
					}
					if q == nil {
						pc.fail = root
					}
				}
				queue = append(queue, pc)
			}
		}
	}
}

// word 是主串
func (root *ACNode) Match(word []rune) []string {
	p := root
	var res []string
	for i, r := range word {
		for _, ok := p.children[r]; !ok && p != root; _, ok = p.children[r] {
			p = p.fail // 失败指针发挥作用的地方
		}
		p = p.children[r]
		if p == nil {
			p = root // 如果没有匹配的，从root重新匹配
		}
		tmp := p
		for tmp != root {
			if tmp.isWord {
				// 输出可以匹配的模式串
				start := i - tmp.length + 1
				res = append(res, string(word[start : i+1]))
			}
			tmp = tmp.fail
		}
	}

	return res
}

// 用 k 来替换 word 中的敏感词， 返回是否有敏感词
func (root *ACNode) Replace(word []rune, k rune) ([]rune, bool) {
	ok := false
	p := root
	for i, r := range word {
		for _, ok := p.children[r]; !ok && p != root; _, ok = p.children[r] {
			p = p.fail // 失败指针发挥作用的地方
		}
		p = p.children[r]
		if p == nil {
			p = root // 如果没有匹配的，从root重新匹配
		}
		tmp := p
		for tmp != root {
			if tmp.isWord {
				// 替换可以匹配的模式串
				start := i - tmp.length + 1
				for j := start; j < i+1; j++ {
					word[j] = k
				}
				ok = true
			}
			tmp = tmp.fail
		}
	}

	return word, ok
}

// 敏感词 word 用 n 个 k 来替换
func (root *ACNode) ReplaceWithN(word []rune, k rune, n int) ([]rune, bool) {
	ok := false
	var res []rune
	kr := make([]rune, n)
	for i := range kr {
		kr[i] = k
	}
	p := root
	prevIndx := 0 // 上一个敏感词的结尾下标
	for i, r := range word {
		for _, ok := p.children[r]; !ok && p != root; _, ok = p.children[r] {
			p = p.fail // 失败指针发挥作用的地方
		}
		p = p.children[r]
		if p == nil {
			p = root // 如果没有匹配的，从root重新匹配
		}
		tmp := p
		for tmp != root {
			if tmp.isWord {
				start := i - tmp.length + 1
				// [prevIndx : start] 为非敏感词
				res = append(res, word[prevIndx: start]...)
				prevIndx = i + 1
				// 屏蔽敏感词
				res = append(res, kr...)
				ok = true
			}
			tmp = tmp.fail
		}
		if i == len(word) - 1 {
			// 添加最后一个非敏感词
			res = append(res, word[prevIndx:]...)
		}
	}

	return res, ok
}