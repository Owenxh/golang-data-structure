package sort

import (
	"testing"
)

func quickSortData() [][]int {
	return [][]int{
		randomIntSliceWithBound(100000, 100000),
		randomIntSliceWithBound(1000000, 1000000),
		randomIntSliceWithBound(2000000, 2000000),
		randomIntSliceWithBound(5000000, 5000000),
		randomIntSliceWithBound(10000000, 10000000),
		randomIntSlice(10000000),
		orderedIntSlice(10000000, 10000000),
		randomIntSliceWithBound(10000000, 1),
	}
}

func TestQuickSort1Way(t *testing.T) {
	data := quickSortData()
	applySortTest(QuickSort1Way, t, data)
}

func TestQuickSort2Ways(t *testing.T) {
	data := quickSortData()
	applySortTest(QuickSort2Ways, t, data)
}

func TestQuickSort3Ways(t *testing.T) {
	data := quickSortData()
	applySortTest(QuickSort3Ways, t, data)
}
