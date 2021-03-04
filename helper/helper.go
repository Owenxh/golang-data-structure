package helper

import (
	"math/rand"
)

func GenerateRandomIntSlice(length, bound int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(bound)
	}
	return res
}

func VerifyIsSorted(arr []int) {
	var prev int
	for i, v := range arr {
		if i > 0 && v < prev {
			panic("Not a sorted slice")
		}
		prev = v
	}
}
