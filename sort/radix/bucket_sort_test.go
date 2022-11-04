package radix

import (
	"io.vava.datastructure/sort"
	"io.vava.datastructure/util"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := util.RandomIntSlice(100_000)

	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	arr3 := make([]int, len(arr))
	copy(arr3, arr)

	util.DoSortTest(func(arr []int) {
		BucketSort(arr, 100)
	}, t, arr)

	util.DoSortTest(func(arr []int) {
		BucketSort(arr, 100)
	}, t, arr2)

	util.DoSortTest(sort.QuickSort2Ways, t, arr3)
}
