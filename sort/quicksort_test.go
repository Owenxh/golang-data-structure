package sort

import (
	"io.vava.datastructure/util"
	"testing"
)

func quickSortData() [][]int {
	return [][]int{
		util.RandomIntSliceWithBound(100000, 100000),
		util.RandomIntSliceWithBound(1000000, 1000000),
		util.RandomIntSliceWithBound(2000000, 2000000),
		util.RandomIntSliceWithBound(5000000, 5000000),
		util.RandomIntSliceWithBound(10000000, 10000000),
		util.RandomIntSlice(10000000),
		util.OrderedIntSlice(10000000, 10000000),
		util.RandomIntSliceWithBound(10000000, 1),
	}
}

func TestQuickSort1Way(t *testing.T) {
	data := quickSortData()
	util.DoSortTests(QuickSort1Way, t, data)
}

func TestQuickSort2Ways(t *testing.T) {
	data := quickSortData()
	util.DoSortTests(QuickSort2Ways, t, data)
}

func TestQuickSort3Ways(t *testing.T) {
	data := quickSortData()
	util.DoSortTests(QuickSort3Ways, t, data)
}
