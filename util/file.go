package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const RootDir = "/golang-data-structure"
const PrideAndPrejudice = "/data/pride-and-prejudice.txt"

func ReadFile(fileName string) (words []string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Print(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return
}

func GetPrideAndPrejudiceWords() []string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = strings.ReplaceAll(path, "\\", "/")

	filePath := path[:strings.Index(path, RootDir)+len(RootDir)+1] + PrideAndPrejudice

	words := ReadFile(filePath)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	fmt.Printf("[Pride-And-Prejudice] words count: %d\n", len(words))
	return words
}

func GetPrideAndPrejudiceAsString() string {
	var sb strings.Builder

	words := GetPrideAndPrejudiceWords()
	for i := 0; i < len(words)-1; i++ {
		sb.WriteString(words[i])
		sb.WriteString(" ")
	}
	sb.WriteString(words[len(words)-1])

	return sb.String()
}

func GetFileAbsolutePath(subPath string) string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = strings.ReplaceAll(path, "\\", "/")

	return path[:strings.Index(path, RootDir)+len(RootDir)+1] + subPath
}
