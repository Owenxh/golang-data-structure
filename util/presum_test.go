package util

import (
	"fmt"
	"testing"
)

func TestPreSum(t *testing.T) {
	nums := []int{-2, 0, 2}

	fmt.Println(buildPreSum(nums))

	fmt.Println(rangeSumCount(nums, 0))
	fmt.Println(rangeSumCount(nums, -2))

	fmt.Println(rangeSumCountBad(nums, 0))
	fmt.Println(rangeSumCountBad(nums, -2))

}
