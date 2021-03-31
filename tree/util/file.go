package util

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var wordReg, _ = regexp.Compile(`[a-zA-Z]+`)

type FileOperation struct {
}

func (f *FileOperation) ReadFile(fileName string) (words []string) {
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
