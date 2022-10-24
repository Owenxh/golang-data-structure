package tree

import (
	"fmt"
	"testing"
)

func TestSegmentTree(t *testing.T) {

	nums := []int{-2, 0, 3, -5, 2, -1}

	st := NewSegmentTree(nums,
		func(l, r int) int {
			return l + r
		})
	fmt.Println(st)
	fmt.Println(st.Query(0, 2))
	fmt.Println(st.Query(2, 5))
	fmt.Println(st.Query(0, 5))

}
