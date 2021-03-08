package sort

import (
	"testing"
)

func TestSelectSort(t *testing.T) {
	applySortTest(SelectSort, t, randomIntSlices(10000, 100000, 200000))
}

func TestSelectSort2(t *testing.T) {
	applySortTest(SelectSort2, t, randomIntSlices(10000, 100000, 200000))
}
