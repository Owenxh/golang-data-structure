package sort

import (
	"testing"
)

func TestMergeSortBU(t *testing.T) {
	data := [][]int{
		randomIntSliceWithBound(5, 5),
		randomIntSliceWithBound(100000, 100000),
		randomIntSliceWithBound(1000000, 1000000),
		randomIntSliceWithBound(2000000, 2000000),
		randomIntSliceWithBound(5000000, 5000000),
	}

	copiedData := make([][]int, len(data), cap(data))
	for i, datum := range data {
		copiedItem := make([]int, len(datum), cap(datum))
		for j, v := range datum {
			copiedItem[j] = v
		}
		copiedData[i] = copiedItem
	}

	applySortTest(MergeSortBU, t, data)
	applySortTest(MergeSort, t, copiedData)
}
