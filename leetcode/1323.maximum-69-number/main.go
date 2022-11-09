package main

import (
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	ret, _ := strconv.Atoi(strings.Replace(str, "6", "9", 1))
	return ret
}
