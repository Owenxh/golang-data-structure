package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := [][]int{
		randomIntSlice(10000, 10000),
		randomIntSlice(100000, 100000),
		randomIntSlice(1000000, 1000000),
		randomIntSlice(2000000, 2000000),
		randomIntSlice(5000000, 5000000),
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
}
