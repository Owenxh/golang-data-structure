package tree

type Trie struct {
	// root 表示 Trie 中的根结点
	root *TrieNode
	// size 表示有 Trie 中有多个单词
	size int
}

type TrieNode struct {
	isWord bool
	next   map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isWord: false,
		next:   make(map[rune]*TrieNode),
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
		size: 0,
	}
}

func (t *Trie) Size() int {
	return t.size
}

// Add 向 Trie 中添加单词 word
func (t *Trie) Add(word string) {
	if word == "" {
		return
	}
	cur := t.root
	for i, runes := 0, []rune(word); i < len(runes); i++ {
		r := runes[i]
		if _, ok := cur.next[r]; !ok {
			cur.next[r] = NewTrieNode()
		}
		cur = cur.next[r]
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

// Add2 use recursion mode
func (t *Trie) Add2(word string) {
	if word == "" {
		return
	}
	t.add(t.root, []rune(word), 0)
}

func (t *Trie) add(root *TrieNode, runes []rune, index int) {
	r := runes[index]

	next, ok := root.next[r]
	if !ok {
		root.next[r] = NewTrieNode()
		next = root.next[r]
	}

	index += 1
	if index == len(runes) {
		if !next.isWord {
			next.isWord = true
			t.size++
			return
		}
	} else {
		t.add(root.next[r], runes, index)
	}
}

// Contains 查询单词 word 是否在 Trie 中
func (t *Trie) Contains(word string) bool {
	if word == "" {
		return false
	}
	cur := t.root
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		next, ok := cur.next[r]
		if !ok {
			return false
		}
		cur = next
	}
	return cur.isWord
}
