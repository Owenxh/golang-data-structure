package util

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

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
