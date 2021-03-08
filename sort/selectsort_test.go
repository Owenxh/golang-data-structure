package sort

import (
	"testing"
	"time"

	"io.vava.datastructure/helper"
)

func internalTestSelectSort(sort func([]int), t *testing.T) {
	arrays := [][]int{
		helper.GenerateRandomIntSlice(1000, 1000),
		helper.GenerateRandomIntSlice(10000, 10000),
		helper.GenerateRandomIntSlice(100000, 100000),
	}

	for _, arr := range arrays {
		start := time.Now()
		sort(arr)
		t.Logf("sort slice(length %d) cost %v", len(arr), time.Now().Sub(start))
		helper.VerifyIsSorted(arr)
	}
}

func TestSelectSort(t *testing.T) {
	internalTestSelectSort(SelectSort, t)
}

func TestSelectSort2(t *testing.T) {
	internalTestSelectSort(SelectSort2, t)
}
