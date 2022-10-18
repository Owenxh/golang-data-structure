package sort

import (
	"io.vava.datastructure/util"
	"testing"
)

func TestMergeSortBU(t *testing.T) {
	data := [][]int{
		util.RandomIntSliceWithBound(5, 5),
		util.RandomIntSliceWithBound(100000, 100000),
		util.RandomIntSliceWithBound(1000000, 1000000),
		util.RandomIntSliceWithBound(2000000, 2000000),
		util.RandomIntSliceWithBound(5000000, 5000000),
	}

	copiedData := make([][]int, len(data), cap(data))
	for i, datum := range data {
		copiedItem := make([]int, len(datum), cap(datum))
		for j, v := range datum {
			copiedItem[j] = v
		}
		copiedData[i] = copiedItem
	}

	util.DoSortTests(MergeSortBU, t, data)
	util.DoSortTests(MergeSort, t, copiedData)
}
