package sort

import "testing"

func TestQuickSort(t *testing.T) {
	data := [][]int{
		randomIntSliceWithBound(100000, 100000),
		randomIntSliceWithBound(1000000, 1000000),
		randomIntSliceWithBound(2000000, 2000000),
		randomIntSliceWithBound(5000000, 5000000),
		randomIntSliceWithBound(10000000, 10000000),
		orderedIntSlice(100000, 100000),
		randomIntSliceWithBound(100000, 1),
	}
	applySortTest(QuickSort, t, data)
}
