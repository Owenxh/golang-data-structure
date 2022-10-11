package tree

import (
	"fmt"
	"io.vava.datastructure/tree/util"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTreePerformance(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := path + "/data/pride-and-prejudice.txt"
	words := util.ReadFile(filePath)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	fmt.Printf("[Pride-And-Prejudice] words count: %d\n", len(words))

	//sort.Slice(words, func(i, j int) bool {
	//	return words[i] < words[j]
	//})

	bst := &BST[string, int]{}
	start := time.Now()
	testBST(bst, words)
	fmt.Printf("[BST] cost time:%v\n", time.Now().Sub(start))
	fmt.Printf("Total different words:%v\n", bst.Size())
	fmt.Printf("Frequency of PRIDE: %v\n", bst.Get("pride"))
	fmt.Printf("Frequency of PREJUDICE: %v\n", bst.Get("prejudice"))

	avl := &AVLTree[string, int]{}
	start = time.Now()
	testAVLTree(avl, words)
	fmt.Printf("[AVL] cost time:%v\n", time.Now().Sub(start))
	fmt.Printf("Total different words:%v\n", avl.Size())
	fmt.Printf("Frequency of PRIDE: %v\n", avl.Get("pride"))
	fmt.Printf("Frequency of PREJUDICE: %v\n", avl.Get("prejudice"))

	if !avl.IsValidBST() {
		t.Fatal("[AVL] isn't valid BST")
	}
	if !avl.IsBalanced() {
		t.Fatal("[AVL] isn't balanced")
	}

	for _, word := range words {
		avl.Remove(word)
		if !avl.IsValidBST() {
			t.Fatal("[AVL] isn't valid BST")
		}
		if !avl.IsBalanced() {
			t.Fatal("[AVL] isn't balanced")
		}
	}

	m := make(map[string]int)
	start = time.Now()
	testMap(m, words)
	fmt.Printf("[MAP] cost time:%v\n", time.Now().Sub(start))
	fmt.Printf("Total different words:%v\n", len(m))
	fmt.Printf("Frequency of PRIDE: %v\n", m["pride"])
	fmt.Printf("Frequency of PREJUDICE: %v\n", m["prejudice"])
}

func testBST(bst *BST[string, int], words []string) {
	for _, word := range words {
		if bst.Contains(word) {
			bst.Add(word, bst.Get(word)+1)
		} else {
			bst.Add(word, 1)
		}
	}
}

func testAVLTree(avl *AVLTree[string, int], words []string) {
	for _, word := range words {
		if avl.Contains(word) {
			avl.Add(word, avl.Get(word)+1)
		} else {
			avl.Add(word, 1)
		}
	}
}

func testMap(m map[string]int, words []string) {
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word] = m[word] + 1
		} else {
			m[word] = 1
		}
	}
}
