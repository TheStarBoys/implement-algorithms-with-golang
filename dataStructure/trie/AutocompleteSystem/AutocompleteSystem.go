package AutocompleteSystem

import (
	"sort"
	"bytes"
)

// todo it has bug
type AutocompleteSystem struct {
	children map[byte]*AutocompleteSystem
	sum int // 总热度，为所有children热度之和
	sentence []byte // 已输入的字符串
	cur *AutocompleteSystem // 最后一个输入字符所对应的节点
}


func Constructor(sentences []string, times []int) AutocompleteSystem {
	trie := constructor()
	for i, sentence := range sentences {
		trie.insert(sentence, times[i])
	}

	return *trie
}

func constructor() *AutocompleteSystem {
	return &AutocompleteSystem{
		children: make(map[byte]*AutocompleteSystem),
	}
}

func (this *AutocompleteSystem) insert(sentence string, times int) {
	trie := this
	for i := range sentence {
		next, ok := trie.children[sentence[i]]
		if !ok {
			next = constructor()
			trie.children[sentence[i]] = next
		}
		next.sum += times
		trie = next
	}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.insert(string(this.sentence), 1)
		this.sentence = []byte{}
		this.cur = nil
		return []string{}
	}
	this.sentence = append(this.sentence, c)
	if this.cur == nil {
		this.cur = this
	}
	next, ok := this.cur.children[c]
	if !ok {
		return []string{}
	}
	this.cur = next
	// 在next的孩子中找到热度前三字符串
	bts := next.findHot()
	return hotest(bts)
}

func hotest(data []*hotSentence) []string {
	sort.Slice(data, func(i, j int) bool {
		if data[i].times > data[j].times {
			return true
		} else {
			if bytes.Compare(data[i].sentence, data[j].sentence) == -1 {
				return true
			}
		}
		return false
	})
	var strs []string
	for _, s := range data {
		strs = append(strs, string(s.sentence))
	}
	if len(strs) < 3 {
		return strs
	}
	return strs[:3]
}

type hotSentence struct {
	sentence []byte
	times int
}

func (this *AutocompleteSystem) findHot() []*hotSentence {
	var res []*hotSentence
	var sentences [][]byte
	stack := make([]*AutocompleteSystem, 0)
	stack = append(stack, this)
	sentences = append(sentences, []byte{})
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		sentence := sentences[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(sentences) != 0 {
			sentences = sentences[:len(sentences)-1]
		}
		if len(cur.children) == 0 {
			res = append(res, &hotSentence{
				sentence: sentence,
				times: cur.sum,
			})
			continue
		}
		for c, next := range cur.children {
			stack = append(stack, next)
			sentence = append(sentence, c)
			sentences = append(sentences, sentence)
		}
	}

	return res
}


/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */
