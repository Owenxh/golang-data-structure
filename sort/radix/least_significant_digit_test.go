package radix

import (
	"fmt"
	"testing"
)

func TestLSDSort(t *testing.T) {
	arr := []string{"BCA", "CAB", "ACB", "BAC", "ABC", "CBA"}
	LSDSort(arr, 3)
	for _, s := range arr {
		fmt.Println(s)
	}
}
