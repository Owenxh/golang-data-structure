package sort

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	data := randomIntSlices(10000, 100000, 5000000)
	copiedData := make([][]int, len(data), cap(data))
	for i, datum := range data {
		copiedItem := make([]int, len(datum), cap(datum))
		for j, v := range datum {
			copiedItem[j] = v
		}
		copiedData[i] = copiedItem
	}

	applySortTest(ShellSort, t, data)
	applySortTest(QuickSort3Ways, t, copiedData)
	if !reflect.DeepEqual(data, copiedData) {
		t.Fatal("invalid sort implementation")
	}
}
