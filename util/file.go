package util

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const PrideAndPrejudice = "/data/pride-and-prejudice.txt"

func ReadFile(fileName string) (words []string) {
	wordReg, err := regexp.Compile(`[a-zA-Z]+`)
	if err != nil {
		panic(err)
	}

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
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words = append(words, wordReg.FindAllString(scanner.Text(), -1)...)
	}
	return
}

func GetPrideAndPrejudiceWords() []string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := path + PrideAndPrejudice
	filePath = strings.ReplaceAll(filePath, "\\", "/")

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
