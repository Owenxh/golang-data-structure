package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := [][]int{
		randomIntSliceWithBound(10000, 10000),
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

	applySortTest(MergeSort, t, data)
	applySortTest(MergeSort2, t, copiedData)

	data = make([][]int, 1)
	data[0] = orderedIntSlice(5000000, 5000000)
	applySortTest(MergeSort2, t, data)
}
