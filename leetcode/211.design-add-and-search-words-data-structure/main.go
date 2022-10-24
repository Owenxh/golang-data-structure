package main

//func main() {
//}

type Node struct {
	isWord bool
	next   map[rune]*Node
}

func NewNode() *Node {
	return &Node{
		next: make(map[rune]*Node),
	}
}

type WordDictionary struct {
	root *Node
	size int
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: NewNode(),
		size: 0,
	}
}

func (w *WordDictionary) AddWord(word string) {
	cur := w.root
	for i, runes := 0, []rune(word); i < len(runes); i++ {
		next, ok := cur.next[runes[i]]
		if !ok {
			cur.next[runes[i]] = NewNode()
			next = cur.next[runes[i]]
		}
		cur = next
	}
	if !cur.isWord {
		cur.isWord = true
		w.size++
	}
}

func (w *WordDictionary) Search(word string) bool {
	//return w.search(w.root, []rune(word), 0)
	return w.match(w.root, []rune(word), 0)
}

func (w *WordDictionary) match(root *Node, runes []rune, index int) bool {
	if index == len(runes) {
		return root.isWord
	}

	r := runes[index]

	if r != '.' {
		if next, ok := root.next[r]; ok {
			return w.match(next, runes, index+1)
		} else {
			return false
		}
	}

	for _, next := range root.next {
		if w.match(next, runes, index+1) {
			return true
		}
	}

	return false
}
