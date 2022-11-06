package main

import (
	"fmt"
	"sort"
	"time"

	"io.vava.datastructure/tree"
	"io.vava.datastructure/util"
)

func main() {
	words := util.GetPrideAndPrejudiceWords()

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	bst := &tree.BST[string, int]{}
	for _, word := range words {
		bst.Add(word, 1)
	}

	start := time.Now()
	for _, word := range words {
		bst.Contains(word)
	}
	fmt.Printf("[BST] cost time:%v\n", time.Now().Sub(start))

	avl := &tree.AVLTree[string, int]{}
	for _, word := range words {
		avl.Add(word, 1)
	}
	fmt.Printf("[AVL] is valid BST: %v\n", avl.IsValidBST())
	fmt.Printf("[AVL] is balanced: %v\n", avl.IsBalanced())

	start = time.Now()
	for _, word := range words {
		avl.Contains(word)
	}
	fmt.Printf("[AVL] cost time:%v\n", time.Now().Sub(start))

	m := make(map[string]int)
	for _, word := range words {
		m[word] = m[word] + 1
	}

	start = time.Now()
	for _, word := range words {
		_, _ = m[word]
	}
	fmt.Printf("[MAP] cost time:%v\n", time.Now().Sub(start))
}
