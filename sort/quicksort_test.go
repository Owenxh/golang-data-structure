package sort

import (
	"testing"
)

func quickSortData() [][]int {
	return [][]int{
		randomIntSliceWithBound(100000, 100000),
		randomIntSliceWithBound(1000000, 1000000),
		randomIntSliceWithBound(2000000, 2000000),
		randomIntSliceWithBound(5000000, 5000000),
		randomIntSliceWithBound(10000000, 10000000),
		randomIntSlice(10000000),
		orderedIntSlice(10000000, 10000000),
		randomIntSliceWithBound(10000000, 1),
	}
}

func TestQuickSort1Way(t *testing.T) {
	data := quickSortData()
	applySortTest(QuickSort1Way, t, data)
}

func TestQuickSort2Ways(t *testing.T) {
	data := quickSortData()
	applySortTest(QuickSort2Ways, t, data)
}

func TestQuickSort3Ways(t *testing.T) {
	data := quickSortData()
	applySortTest(QuickSort3Ways, t, data)
}

func TestInternalQuickSort3Ways(t *testing.T) {
	data := quickSortData()
	applySortTest(InternalQuickSort3Ways, t, data)
}

func InternalQuickSort3Ways(arr []int) {
	internalQuickSort3Ways(arr, 0, len(arr)-1)
}

func internalQuickSort3Ways(arr []int, l, r int) {
	if l >= r {
		return
	}
	lt, gt := internalPartition3Ways(arr, l, r)
	internalQuickSort3Ways(arr, l, lt-1)
	internalQuickSort3Ways(arr, gt, r)
}

// 循环不变量
// ① [l+1, lt] < v
// ② [lt + 1, i - 1] == v
// ③ [gt, r] > v
func internalPartition3Ways(arr []int, l, r int) (int, int) {
	v := arr[l]
	lt, gt := l, r+1

	for i := l + 1; i < gt; {
		if arr[i] < v {
			lt++
			arr[lt], arr[i] = arr[i], arr[lt]
			i++
		} else if arr[i] > v {
			gt--
			arr[gt], arr[i] = arr[i], arr[gt]
		} else {
			i++
		}
	}
	arr[lt], arr[l] = arr[l], arr[lt]
	return lt, gt
}
