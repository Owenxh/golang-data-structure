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

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := path + "/tree/data/pride-and-prejudice.txt"
	words := fileOperation.ReadFile(filePath)
	bst := &tree.BST{}

	start := time.Now()
	for _, word := range words {
		bst.Add(strings.ToLower(word))
	}
	for _, word := range words {
		if ok := bst.Contains(strings.ToLower(word)); !ok {
			panic("error implementation of BST")
		}
	}

	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
	fmt.Printf("[pride-and-prejudice] words count: %d\n", len(words))
	fmt.Printf("[pride-and-prejudice] different words count: %d\n", bst.Size())

	start = time.Now()
	theMap := make(map[string]int)
	for _, word := range words {
		theMap[strings.ToLower(word)] = theMap[strings.ToLower(word)] + 1
	}
	for _, word := range words {
		if theMap[strings.ToLower(word)] == 0 {
			panic("error implementation of BST")
		}
	}

	fmt.Printf("cost time:%v\n", time.Now().Sub(start))
	fmt.Printf("[pride-and-prejudice] words count: %d\n", len(words))
	fmt.Printf("[pride-and-prejudice] different words count: %d\n", bst.Size())
}
