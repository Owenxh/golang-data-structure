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

func TestLazySegmentTree(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	st := NewLazySegmentTree(nums)

	var l, r, diff int

	l, r = 0, 2
	fmt.Printf("[%v, %v] = %v\n", l, r, st.Query(l, r))

	l, r = 3, 5
	fmt.Printf("[%v, %v] = %v\n", l, r, st.Query(l, r))

	l, r = 0, 5
	fmt.Printf("[%v, %v] = %v\n", l, r, st.Query(l, r))

	l, r, diff = 0, 5, 5
	fmt.Printf("update range [%v, %v] = %v\n", l, r, diff)
	st.UpdateRange(l, r, diff)

	l, r = 0, 2
	fmt.Printf("[%v, %v] = %v\n", l, r, st.Query(l, r))

	l, r, diff = 0, 4, 5
	fmt.Printf("update range [%v, %v] = %v\n", l, r, diff)
	st.UpdateRange(l, r, diff)

	l, r = 3, 5
	fmt.Printf("[%v, %v] = %v\n", l, r, st.Query(l, r))

	l, r = 0, 5
	fmt.Printf("[%v, %v] = %v\n", l, r, st.Query(l, r))
}
