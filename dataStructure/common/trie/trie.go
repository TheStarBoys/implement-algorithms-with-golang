package trie

type Trie interface {
	Insert(word []rune) // Insert insert word into trie, returns the node at the last char of the word
	Search(word []rune) bool // Search returns true if the word exist in trie
}

// Children is Trie children node
type Children interface {
	Get(r rune) (Trie, bool)
	Put(r rune, trie Trie)
	Del(r rune)
	Len() int // Len returns number of Children, not data size
	New() Children // New returns new children
}

type TrieImpl struct {
	children Children
	isWord bool
}

func New(children Children) Trie {
	return &TrieImpl{
		children: children,
	}
}

func (t *TrieImpl) Insert(word []rune) {
	trie := t
	for _, r := range word {
		next, ok := trie.children.Get(r)
		if !ok {
			tmp := New(t.children.New())
			next = tmp
			trie.children.Put(r, next)
		}
		trie = next.(*TrieImpl)
	}
	trie.isWord = true
}

func (t *TrieImpl) Search(word []rune) bool {
	trie := t
	for _, r := range word {
		next, ok := trie.children.Get(r)
		if !ok {
			return false
		}
		trie = next.(*TrieImpl)
	}

	return trie.isWord || trie.children.Len() == 0
}