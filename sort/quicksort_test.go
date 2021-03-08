package sort

import "testing"

func TestQuickSort(t *testing.T) {
	data := [][]int{
		randomIntSlice(10000, 10000),
		randomIntSlice(100000, 100000),
		randomIntSlice(1000000, 1000000),
		randomIntSlice(2000000, 2000000),
		randomIntSlice(5000000, 5000000),
		randomIntSlice(10000000, 10000000),
	}
	applySortTest(QuickSort, t, data)
}
