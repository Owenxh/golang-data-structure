package sort

import (
	"testing"
)

func selectSortData() [][]int {
	return [][]int{
		randomIntSliceWithBound(100000, 100000),
		// randomIntSliceWithBound(1000000, 1000000),
		// randomIntSliceWithBound(2000000, 2000000),
		// randomIntSliceWithBound(5000000, 5000000),
		// randomIntSliceWithBound(10000000, 10000000),
		// randomIntSlice(10000000),
		// orderedIntSlice(10000000, 10000000),
		// randomIntSliceWithBound(10000000, 1),
	}
}

func TestSelectSort(t *testing.T) {
	data := selectSortData()
	applySortTest(SelectSort, t, data)
}
