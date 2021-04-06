package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"io.vava.datastructure/tree"

	"io.vava.datastructure/tree/util"
)

var fileOperation = &util.FileOperation{}

type operation interface {
	add(string)
	remove(string)
	contains(string) bool
	size() int
}

type AVLTreeOperation struct {
	t *tree.AVLTree
}

func (o *AVLTreeOperation) add(s string) {
	o.t.Add(s)
}

func (o *AVLTreeOperation) remove(s string) {
	o.t.Remove(s)
}

func (o *AVLTreeOperation) contains(s string) bool {
	return o.t.Contains(s)
}

func (o *AVLTreeOperation) size() int {
	return o.t.Size()
}

type BSTOperation struct {
	t *tree.BST
}

func (o *BSTOperation) add(s string) {
	o.t.Add(s)
}

func (o *BSTOperation) remove(s string) {
	o.t.Remove(s)
}

func (o *BSTOperation) contains(s string) bool {
	return o.t.Contains(s)
}

func (o *BSTOperation) size() int {
	return o.t.Size()
}

type MapOperation struct {
	m map[string]int
}

func (o *MapOperation) add(s string) {
	o.m[s] = o.m[s] + 1
}

func (o *MapOperation) remove(s string) {
	delete(o.m, s)
}

func (o *MapOperation) contains(s string) bool {
	return o.m[s] > 0
}

func (o *MapOperation) size() int {
	return len(o.m)
}

func test(words []string, fn operation) (count int) {
	for _, word := range words {
		fn.add(strings.ToLower(word))
	}
	for _, word := range words {
		if ok := fn.contains(strings.ToLower(word)); !ok {
			panic("error implementation of BST")
		}
	}
	count = fn.size()
	for _, word := range words {
		fn.remove(strings.ToLower(word))
	}

	if fn.size() != 0 {
		panic("error implementation of BST")
	}
	return
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := path + "/tree/data/pride-and-prejudice.txt"
	words := fileOperation.ReadFile(filePath)
	fmt.Printf("[Pride-And-Prejudice] words count: %d\n", len(words))

	start := time.Now()
	count := test(words, &BSTOperation{&tree.BST{}})
	fmt.Printf("[BST] cost time:%v\n", time.Now().Sub(start))
	fmt.Printf("[BST] different words count: %d\n", count)

	start = time.Now()
	count = test(words, &AVLTreeOperation{&tree.AVLTree{}})
	fmt.Printf("[AVL Tree] cost time:%v\n", time.Now().Sub(start))
	fmt.Printf("[AVL Tree] different words count: %d\n", count)

	start = time.Now()
	count = test(words, &MapOperation{make(map[string]int)})
	fmt.Printf("[MAP] cost time:%v\n", time.Now().Sub(start))
	fmt.Printf("[MAP] different words count: %d\n", count)
}
