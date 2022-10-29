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

func execQuery(st *LazySegmentTree, l, r int) {
	fmt.Printf(" Query range [%v, %v] = %v\n", l, r, st.Query(l, r))
}

func execUpdate(st *LazySegmentTree, l, r, diff int) {
	fmt.Printf("Update range [%v, %v] = %v\n", l, r, diff)
	st.UpdateRange(l, r, diff)
}

func TestLazySegmentTree(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	st := NewLazySegmentTree(nums)

	execQuery(st, 0, 2)
	execQuery(st, 3, 5)
	execQuery(st, 0, 5)

	execUpdate(st, 0, 5, 5)
	execQuery(st, 0, 5)
	execQuery(st, 0, 2)

	execUpdate(st, 0, 4, 5)
	execUpdate(st, 3, 4, 4)
	execQuery(st, 3, 5)
	execQuery(st, 0, 5)
	execQuery(st, 2, 3)
	execQuery(st, 0, 2)
	execQuery(st, 3, 4)
	execQuery(st, 5, 5)

	st.RankVisit(func(index, v int, l, r int) {
		fmt.Printf("Node %d range [%v, %v] = %v\n", index, l, r, v)
	})

	fmt.Println("reversed rank visit")
	st.ReversedRankVisit(func(index, v int, l, r int) {
		fmt.Printf("Node %d range [%v, %v] = %v\n", index, l, r, v)
	})

}
