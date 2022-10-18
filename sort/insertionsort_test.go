package sort

import (
	"io.vava.datastructure/util"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	util.DoSortTests(InsertionSort, t, util.RandomIntSlices(10000, 100000, 200000))
}

func TestInsertionSort2(t *testing.T) {
	util.DoSortTests(InsertionSort2, t, util.RandomIntSlices(10000, 100000, 200000))
}
