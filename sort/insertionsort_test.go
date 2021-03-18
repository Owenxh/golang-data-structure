package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	applySortTest(InsertionSort, t, randomIntSlices(10000, 100000, 200000))
}

func TestInsertionSort2(t *testing.T) {
	applySortTest(InsertionSort2, t, randomIntSlices(10000, 100000, 200000))
}
