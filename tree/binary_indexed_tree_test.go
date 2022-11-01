package tree

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBinaryIndexedTree(t *testing.T) {

	for i := 0; i < 10; i++ {
		fmt.Printf("%08v %08v	\n", strconv.FormatInt(int64(i), 2), strconv.FormatInt(int64(lowBit(i)), 2))
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	bit := NewBIT(nums)

	// [0 1 3 3 10 5 11 7 36]

	fmt.Println(bit.SumRange(0, 7))
	fmt.Println(bit.SumRange(0, 7))
	fmt.Println(bit.SumRange(4, 6))

	bit.Update(5, 10)

	fmt.Println(bit.SumRange(0, 7))
	fmt.Println(bit.SumRange(0, 7))
	fmt.Println(bit.SumRange(4, 6))
	fmt.Println(bit.SumRange(0, 4))
}
