package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"io.vava.datastructure/tree"

	"io.vava.datastructure/tree/util"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := path + "/tree/data/pride-and-prejudice.txt"
	words := util.ReadFile(filePath)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	fmt.Printf("[Pride-And-Prejudice] words count: %d\n", len(words))

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	bst := &tree.BST{}
	for _, word := range words {
		bst.Add(word)
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
