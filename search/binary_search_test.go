package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i <= len(arr); i++ {
		fmt.Print(BinarySearch(arr, i), " ")
	}
	fmt.Println()
}
