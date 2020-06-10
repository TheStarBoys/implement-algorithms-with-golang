package trie

import (
	"testing"
	"fmt"
)

// mapChildren
type mapChildren map[rune]Trie

func NewMapChildren() *mapChildren {
	m :=  mapChildren(map[rune]Trie{})
	return &m
}

func (m mapChildren) Get(r rune) (Trie, bool) {
	t, ok :=  m[r]
	return t, ok
}

func (m *mapChildren) Put(r rune, t Trie) {
	(*m)[r] = t
}

func (m *mapChildren) Del(r rune) {
	delete(*m, r)
}

func (m mapChildren) Len() int {
	return len(m)
}

func (m mapChildren) New() Children {
	return NewMapChildren()
}

// Slice children
type sliceChildren []*sliceChildrenVal

type sliceChildrenVal struct {
	r rune
	t Trie
}

func NewSliceChildren() *sliceChildren {
	i := make([]*sliceChildrenVal, 0, 4)
	s := sliceChildren(i)
	return &s
}

func (s sliceChildren) Get(r rune) (Trie, bool) {
	l, h := 0, s.Len() - 1
	for l <= h {
		mid := l + ((h - l) >> 1)
		if s[mid].r == r {
			return s[mid].t, true
		} else if s[mid].r < r {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}

	return nil, false
}

func (s *sliceChildren) Put(r rune, t Trie) {
	if s.Len() == 0 {
		*s = append(*s, &sliceChildrenVal{r, t})
		return
	}
	i := s.Len() - 1
	*s = append(*s, &sliceChildrenVal{})
	for i >= 0 && (*s)[i].r > r {
		(*s)[i+1] = (*s)[i]
		i--
	}
	(*s)[i+1] = &sliceChildrenVal{r, t}
}

func (s *sliceChildren) Del(r rune) {
	l, h := 0, s.Len() - 1
	for l <= h {
		mid := l + ((h - l) >> 1)
		if (*s)[mid].r == r {
			*s = append((*s)[:mid], (*s)[mid+1:]...)
			return
		} else if (*s)[mid].r < r {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
}

func (s *sliceChildren) Len() int {
	return len(*s)
}

func (s *sliceChildren) New() Children {
	return NewSliceChildren()
}

func (s *sliceChildren) String() string {
	res := "sliceChildren ["
	for i, v := range *s {
		if i == s.Len() - 1 {
			res += fmt.Sprintf("%d:%v]", v.r, v.t)
			break
		}
		res += fmt.Sprintf("%d:%v ", v.r, v.t)
	}

	return res
}

func TestTrie(t *testing.T) {
	// trie.Children 被定义为一个接口，其目的是可以轻松的切换trie孩子节点的具体实现
	// 根据具体的项目需求，可能并不需要用数组（或者切片）或者散列表来达到 O(1)的查询效率，因为其比较浪费空间
	// 此文件还提供了一个有序切片，减少空间复杂度，通过二分查找来达到 O(log n)的查询效率（时间复杂度提高了），n 为孩子节点的个数
	// 同理，children可以被实现为一个链表，跳表，红黑树等数据结构

	f := func(trie Trie) {
		trie.Insert([]rune("he"))
		trie.Insert([]rune("hello"))
		trie.Insert([]rune("hello,world"))
		trie.Insert([]rune("hello,wor"))
		trie.Insert([]rune("hallo"))
		trie.Insert([]rune("hero"))
		trie.Insert([]rune("zero"))
		trie.Insert([]rune("zone"))
		trie.Insert([]rune("zoom"))
		trie.Insert([]rune("Alice"))
		trie.Insert([]rune("abandon"))
		trie.Insert([]rune("giant"))
		trie.Insert([]rune("great"))
		trie.Insert([]rune("grunt"))

		trie.Insert([]rune("money"))
		trie.Insert([]rune("monkey"))
		trie.Insert([]rune("你好，世界"))
		trie.Insert([]rune("你好"))


		var (
			word string
			b bool
		)
		word = "he"
		b = trie.Search([]rune(word))
		if b == false {
			t.Errorf("search word: %s, expect: true, got: false", word)
		}

		word = "hello,wor"
		b = trie.Search([]rune(word))
		if b == false {
			t.Errorf("search word: %s, expect: true, got: false", word)
		}

		word = "hello,world"
		b = trie.Search([]rune(word))
		if b == false {
			t.Errorf("search word: %s, expect: true, got: false", word)
		}

		word = "abandon"
		b = trie.Search([]rune(word))
		if b == false {
			t.Errorf("search word: %s, expect: true, got: false", word)
		}

		word = "money"
		b = trie.Search([]rune(word))
		if b == false {
			t.Errorf("search word: %s, expect: true, got: false", word)
		}

		word = "zero"
		b = trie.Search([]rune(word))
		if b == false {
			t.Errorf("search word: %s, expect: true, got: false", word)
		}

		word = "你好"
		b = trie.Search([]rune(word))
		if b == false {
			t.Errorf("search word: %s, expect: true, got: false", word)
		}

		word = "你好，世界"
		b = trie.Search([]rune(word))
		if b == false {
			t.Errorf("search word: %s, expect: true, got: false", word)
		}

		word = "Alic"
		b = trie.Search([]rune(word))
		if b == true {
			t.Errorf("search word: %s, expect: false, got: true", word)
		}

		word = "gruntte"
		b = trie.Search([]rune(word))
		if b == true {
			t.Errorf("search word: %s, expect: false, got: true", word)
		}
	}

	trie := New(NewMapChildren())
	f(trie)

	trie = New(NewSliceChildren())
	f(trie)
}
