package util

import (
	"fmt"
	"testing"
)

func TestPreSum(t *testing.T) {
	nums := []int{-2, 0, 2}

	fmt.Println(BuildPreSum(nums))

	fmt.Println(RangeSumCount(nums, 0))
	fmt.Println(RangeSumCount(nums, -2))

	fmt.Println(RangeSumCountBad(nums, 0))
	fmt.Println(RangeSumCountBad(nums, -2))
}
